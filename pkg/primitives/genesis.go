package primitives

import (
	"time"

	"github.com/olympus-protocol/ogen/pkg/chainhash"
	"github.com/olympus-protocol/ogen/pkg/params"
)

// GetGenesisBlock gets the genesis block for a certain chain parameters.
func GetGenesisBlock(params params.ChainParams) Block {
	return Block{
		Header: &BlockHeader{
			Version:                    0,
			Nonce:                      0,
			TxMerkleRoot:               chainhash.Hash{},
			VoteMerkleRoot:             chainhash.Hash{},
			DepositMerkleRoot:          chainhash.Hash{},
			ExitMerkleRoot:             chainhash.Hash{},
			VoteSlashingMerkleRoot:     chainhash.Hash{},
			RANDAOSlashingMerkleRoot:   chainhash.Hash{},
			ProposerSlashingMerkleRoot: chainhash.Hash{},
			GovernanceVotesMerkleRoot:  chainhash.Hash{},
			PrevBlockHash:              chainhash.Hash{},
			Timestamp:                  uint64(time.Unix(0x0, 0).Unix()),
			Slot:                       0,
			StateRoot:                  chainhash.Hash{},
			FeeAddress:                 [20]byte{},
		},
		Txs: []*Tx{},
	}
}