package chain

import (
	"errors"
	"fmt"
	"github.com/olympus-protocol/ogen/internal/txindex"
	"time"

	"github.com/olympus-protocol/ogen/internal/blockdb"
	"github.com/olympus-protocol/ogen/internal/chainindex"
	"github.com/olympus-protocol/ogen/pkg/bls"
	"github.com/olympus-protocol/ogen/pkg/chainhash"
	"github.com/olympus-protocol/ogen/pkg/primitives"
)

type blockRowAndValidator struct {
	row       *chainindex.BlockRow
	validator uint64
}

// UpdateChainHead updates the blockchain head if needed
func (ch *blockchain) UpdateChainHead(txn blockdb.DBUpdateTransaction, possible chainhash.Hash) error {
	_, justifiedState := ch.state.GetJustifiedHead()

	activeValidatorIndices := justifiedState.GetValidatorIndicesActiveAt(justifiedState.GetEpochIndex())
	var targets []blockRowAndValidator
	for _, i := range activeValidatorIndices {
		bl, err := ch.getLatestAttestationTarget(i)
		if err != nil {
			continue
		}
		targets = append(targets, blockRowAndValidator{
			row:       bl,
			validator: i})
	}

	getVoteCount := func(block *chainindex.BlockRow) uint64 {
		votes := uint64(0)
		for _, target := range targets {
			node := target.row.GetAncestorAtSlot(block.Slot)
			if node == nil {
				return 0
			}
			if node.Hash.IsEqual(&block.Hash) {
				votes += justifiedState.GetEffectiveBalance(target.validator, &ch.params) / 1e8
			}
		}
		return votes
	}

	head, _ := ch.state.GetJustifiedHead()

	// this may seem weird, but it occurs when importing when the justified block is not
	// imported, but the finalized head is. It should never occur other than that
	if head == nil {
		head, _ = ch.state.GetFinalizedHead()
	}

	for {
		children := head.Children()
		if len(children) == 0 {
			if head.Hash.IsEqual(&possible) {
				ch.state.Blockchain().SetTip(head)

				ch.log.Infof("setting head to %s", head.Hash)

				err := txn.SetTip(head.Hash)
				if err != nil {
					return err
				}

			}
			return nil
		}
		bestVoteCountChild := children[0]
		bestVotes := uint64(0)
		if len(children) > 1 {
			bestVotes = getVoteCount(bestVoteCountChild)
		}
		for _, c := range children[1:] {
			vc := getVoteCount(c)
			if vc > bestVotes {
				bestVoteCountChild = c
				bestVotes = vc
			}
		}
		head = bestVoteCountChild
	}
}

func (ch *blockchain) getLatestAttestationTarget(validator uint64) (row *chainindex.BlockRow, err error) {
	var att *primitives.MultiValidatorVote
	att, ok := ch.state.GetLatestVote(validator)
	if !ok {
		return nil, fmt.Errorf("attestation target not found")
	}

	row, ok = ch.state.BlockIndex().Get(att.Data.BeaconBlockHash)
	if !ok {
		return nil, errors.New("couldn't find block attested to by validator in chainindex")
	}
	return row, nil
}

// ProcessBlock processes an incoming block from a peer or the miner.
func (ch *blockchain) ProcessBlock(block *primitives.Block) error {
	// 1. first verify basic block properties
	// b. get parent block
	blockTime := ch.genesisTime.Add(time.Second * time.Duration(ch.params.SlotDuration*block.Header.Slot))

	blockHash := block.Hash()

	if other, found := ch.State().Chain().GetNodeBySlot(block.Header.Slot); found && other.Slot == block.Header.Slot {
		if other.Hash.IsEqual(&blockHash) {
			return nil
		}

		lastBlockHash := block.Header.PrevBlockHash

		view, err := ch.State().GetSubView(lastBlockHash)
		if err != nil {
			return err
		}

		lastBlockState, _, err := ch.State().GetStateForHashAtSlot(lastBlockHash, block.Header.Slot, &view, &ch.params)
		if err != nil {
			return err
		}

		if err := lastBlockState.CheckBlockSignature(block, &ch.params); err != nil {
			return err
		}

		otherBlock, err := ch.GetBlock(other.Hash)
		if err != nil {
			return err
		}

		proposerPub, err := lastBlockState.GetProposerPublicKey(block, &ch.params)
		if err != nil {
			return err
		}

		blockSig, err := bls.SignatureFromBytes(block.Signature[:])
		if err != nil {
			return err
		}

		otherSig, err := bls.SignatureFromBytes(otherBlock.Signature[:])
		if err != nil {
			return err
		}

		ch.log.Warnf("found duplicate block at slot %d, reporting...", block.Header.Slot)

		for n := range ch.notifees {
			var b, os [96]byte
			var p [48]byte
			copy(b[:], blockSig.Marshal())
			copy(os[:], otherSig.Marshal())
			copy(p[:], proposerPub.Marshal())
			n.ProposerSlashingConditionViolated(&primitives.ProposerSlashing{
				BlockHeader1:       block.Header,
				BlockHeader2:       otherBlock.Header,
				Signature1:         b,
				Signature2:         os,
				ValidatorPublicKey: p,
			})
		}
		return nil
	}

	if time.Now().Add(time.Second * 2).Before(blockTime) {
		return fmt.Errorf("block %d processed at %s, but should wait until %s", block.Header.Slot, time.Now(), blockTime)
	}

	// 2. verify block against previous block's state

	newState, receipts, err := ch.State().Add(block)
	if err != nil {
		ch.log.Warn(err)
		return err
	}

	if len(receipts) > 0 {
		msg := "\nEpoch Receipts\n----------\n"
		receiptTypes := make(map[string]int64)

		for _, r := range receipts {
			if _, ok := receiptTypes[r.TypeString()]; !ok {
				receiptTypes[r.TypeString()] = r.Amount
			} else {
				receiptTypes[r.TypeString()] += r.Amount
			}
		}

		for rt, amount := range receiptTypes {
			if amount > 0 {
				msg += fmt.Sprintf("rewarded %d for %s\n", amount, rt)
			} else if amount < 0 {
				msg += fmt.Sprintf("penalized %d for %s\n", -amount, rt)
			} else {
				msg += fmt.Sprintf("neutral increments for %s\n", rt)
			}
		}

		ch.log.Debugf(msg)
	}

	return ch.db.Update(func(txn blockdb.DBUpdateTransaction) error {
		err = txn.AddRawBlock(block)
		if err != nil {
			return err
		}

		row, err := ch.state.BlockIndex().Add(*block)
		if err != nil {
			return err
		}

		// set current block row in database
		if err := txn.SetBlockRow(row.ToBlockNodeDisk()); err != nil {
			return err
		}

		// update parent to point at current
		if err := txn.SetBlockRow(row.Parent.ToBlockNodeDisk()); err != nil {
			return err
		}

		for _, a := range block.Votes {
			validators, err := newState.GetVoteCommittee(a.Data.Slot, &ch.params)
			if err != nil {
				return err
			}

			ch.state.SetLatestVotesIfNeeded(validators, a)
		}

		// TODO: remove when we have fork choice
		if err := ch.UpdateChainHead(txn, blockHash); err != nil {
			return err
		}

		view, err := ch.State().GetSubView(block.Header.PrevBlockHash)
		if err != nil {
			return err
		}

		finalizedSlot := newState.GetFinalizedEpoch() * ch.params.EpochLength
		finalizedHash, err := view.GetHashBySlot(finalizedSlot)
		if err != nil {
			return err
		}
		finalizedState, found := ch.state.GetStateForHash(finalizedHash)
		if !found {
			return fmt.Errorf("could not find finalized state with hash %s in state map", finalizedHash)
		}
		if err := txn.SetFinalizedHead(finalizedHash); err != nil {
			return err
		}
		if err := ch.state.setFinalizedHead(finalizedHash, finalizedState); err != nil {
			return err
		}
		if err := txn.SetFinalizedState(finalizedState); err != nil {
			return err
		}

		ch.state.RemoveBeforeSlot(newState.GetFinalizedEpoch() * ch.params.EpochLength)

		justifiedState, found := ch.state.GetStateForHash(newState.GetJustifiedEpochHash())
		if !found {
			return fmt.Errorf("could not find justified state with hash %s in state map", newState.GetJustifiedEpochHash())
		}
		if err := txn.SetJustifiedHead(newState.GetJustifiedEpochHash()); err != nil {
			return err
		}
		if err := ch.state.setJustifiedHead(newState.GetJustifiedEpochHash(), justifiedState); err != nil {
			return err
		}
		if err := txn.SetJustifiedState(justifiedState); err != nil {
			return err
		}

		// TODO: delete state before finalized

		ch.log.Debugf("processed %d votes %d deposits %d exits and %d transactions", len(block.Votes), len(block.Deposits), len(block.Exits), len(block.Txs))
		ch.log.Debugf("included %d vote slashing %d randao slashing %d proposer slashing", len(block.VoteSlashings), len(block.RANDAOSlashings), len(block.ProposerSlashings))
		ch.log.Infof("new block at slot: %d with %d finalized and %d justified", block.Header.Slot, newState.GetFinalizedEpoch(), newState.GetJustifiedEpoch())

		// TODO: add a log that shows network participation with expected.

		// Once a block is accepted build tx chainindex and account tx tracking

		for i, tx := range block.Txs {
			locator := txindex.TxLocator{
				Hash:  tx.Hash(),
				Block: block.Hash(),
				Index: uint64(i),
			}
			from, err := tx.FromPubkeyHash()
			if err != nil {
				return err
			}

			// Add chainindex to senders
			err = ch.txidx.SetTx(locator, from)
			if err != nil {
				return err
			}
			// Add chainindex to receivers
			err = ch.txidx.SetTx(locator, tx.To)
			if err != nil {
				return err
			}
		}
		ch.notifeeLock.RLock()
		stateCopy := newState.Copy()
		for i := range ch.notifees {
			go i.NewTip(row, block, stateCopy, receipts)
		}
		ch.notifeeLock.RUnlock()
		return nil
	})
}
