package blockdb

import (
	"errors"
	"fmt"
	"github.com/olympus-protocol/ogen/internal/state"
	"path"
	"sync"
	"sync/atomic"
	"time"

	"github.com/olympus-protocol/ogen/internal/logger"
	"github.com/olympus-protocol/ogen/pkg/chainhash"
	"github.com/olympus-protocol/ogen/pkg/params"
	"github.com/olympus-protocol/ogen/pkg/primitives"
	"go.etcd.io/bbolt"
)

// BlockDBBucketKey is the bucket key of the blocks on the database.
var BlockDBBucketKey = []byte("blocksdb")

// BlockDB is an interface for blockDb
type BlockDB interface {
	Close()
	Update(cb func(txn DBUpdateTransaction) error) error
	View(cb func(txn DBViewTransaction) error) error
}

var _ BlockDB = &blockDB{}

// BlockDB is the struct wrapper for the block database.
type blockDB struct {
	log    logger.Logger
	db     *bbolt.DB
	params params.ChainParams
	lock   sync.RWMutex

	requestedClose uint32
	canClose       sync.WaitGroup
}

// BlockDBUpdateTransaction is a wrapper for the bbolt transaction with writing privileges.
type BlockDBUpdateTransaction struct {
	BlockDBReadTransaction
}

// BlockDBReadTransaction is a wrapper for the bbolt transaction with view privileges.
type BlockDBReadTransaction struct {
	db  BlockDB
	bkt *bbolt.Bucket
}

// NewBlockDB returns a database instance with a rawBlockDatabase and BboltDB to use on the selected path.
func NewBlockDB(pathDir string, params params.ChainParams, log logger.Logger) (BlockDB, error) {
	db, err := bbolt.Open(path.Join(pathDir, "chain.db"), 0600, nil)
	if err != nil {
		return nil, err
	}
	blockdb := &blockDB{
		log:    log,
		db:     db,
		params: params,
	}
	return blockdb, nil
}

// Close closes the database.
func (bdb *blockDB) Close() {
	if atomic.LoadUint32(&bdb.requestedClose) != 0 {
		return
	}
	atomic.StoreUint32(&bdb.requestedClose, 1)
	bdb.canClose.Wait()
	_ = bdb.db.Close()
}

// Update gets a transaction for updating the database.
func (bdb *blockDB) Update(cb func(txn DBUpdateTransaction) error) error {
	bdb.lock.Lock()
	defer bdb.lock.Unlock()
	if atomic.LoadUint32(&bdb.requestedClose) != 0 {
		return fmt.Errorf("database is closing")
	}

	bdb.canClose.Add(1)
	defer bdb.canClose.Done()

	return bdb.db.Update(func(tx *bbolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists(BlockDBBucketKey)
		if err != nil {
			return err
		}
		blockTxn := BlockDBReadTransaction{
			db:  bdb,
			bkt: bkt,
		}

		return cb(&BlockDBUpdateTransaction{blockTxn})
	})
}

// View gets a transaction for viewing the database.
func (bdb *blockDB) View(cb func(txn DBViewTransaction) error) error {
	bdb.lock.RLock()
	defer bdb.lock.RUnlock()
	if atomic.LoadUint32(&bdb.requestedClose) != 0 {
		return fmt.Errorf("database is closing")
	}

	bdb.canClose.Add(1)
	defer bdb.canClose.Done()
	return bdb.db.View(func(tx *bbolt.Tx) error {
		blockTxn := &BlockDBReadTransaction{
			db:  bdb,
			bkt: tx.Bucket(BlockDBBucketKey),
		}

		return cb(blockTxn)
	})
}

// GetBlock gets a block from the database.
func (brt *BlockDBReadTransaction) GetBlock(hash chainhash.Hash) (*primitives.Block, error) {
	blockBytes, err := getKey(brt, hash[:])
	if err != nil {
		return nil, err
	}

	block := new(primitives.Block)
	err = block.Unmarshal(blockBytes)
	return block, err
}

// GetRawBlock gets a block serialized from the database.
func (brt *BlockDBReadTransaction) GetRawBlock(hash chainhash.Hash) ([]byte, error) {
	blockBytes, err := getKey(brt, hash[:])
	if err != nil {
		return nil, err
	}
	return blockBytes, err
}

// AddRawBlock adds a raw block to the database.
func (but *BlockDBUpdateTransaction) AddRawBlock(block *primitives.Block) error {
	blockHash := block.Hash()
	blockBytes, err := block.Marshal()
	if err != nil {
		return err
	}
	return setKey(but, blockHash[:], blockBytes)
}

func getKeyHash(tx *BlockDBReadTransaction, key []byte) (chainhash.Hash, error) {
	var out chainhash.Hash
	i := tx.bkt.Get(key)
	if len(i) <= 0 {
		return chainhash.Hash{}, errors.New("no data")
	}
	copy(out[:], i)
	return out, nil
}

func getKey(tx *BlockDBReadTransaction, key []byte) ([]byte, error) {
	i := tx.bkt.Get(key)
	if len(i) <= 0 {
		return nil, errors.New("no data")
	}
	return i, nil
}

func setKeyHash(tx *BlockDBUpdateTransaction, key []byte, to chainhash.Hash) error {
	return tx.bkt.Put(key, to[:])
}

func setKey(tx *BlockDBUpdateTransaction, key []byte, to []byte) error {
	return tx.bkt.Put(key, to)
}

var tipKey = []byte("chain-tip")

// SetTip sets the current best tip of the blockchain.
func (but *BlockDBUpdateTransaction) SetTip(c chainhash.Hash) error {
	return setKeyHash(but, tipKey, c)
}

// GetTip gets the current best tip of the blockchain.
func (brt *BlockDBReadTransaction) GetTip() (chainhash.Hash, error) {
	return getKeyHash(brt, tipKey)
}

var finalizedStateKey = []byte("finalized-state")

// SetFinalizedState sets the finalized state of the blockchain.
func (but *BlockDBUpdateTransaction) SetFinalizedState(s state.State) error {
	buf, err := s.Marshal()
	if err != nil {
		return err
	}

	return setKey(but, finalizedStateKey, buf)
}

// GetFinalizedState gets the finalized state of the blockchain.
func (brt *BlockDBReadTransaction) GetFinalizedState() (state.State, error) {
	stateBytes, err := getKey(brt, finalizedStateKey)
	if err != nil {
		return nil, err
	}
	s := state.NewEmptyState()
	err = s.Unmarshal(stateBytes)
	return s, err
}

var justifiedStateKey = []byte("justified-state")

// SetJustifiedState sets the justified state of the blockchain.
func (but *BlockDBUpdateTransaction) SetJustifiedState(s state.State) error {
	buf, err := s.Marshal()
	if err != nil {
		return err
	}

	return setKey(but, justifiedStateKey, buf)
}

// GetJustifiedState gets the justified state of the blockchain.
func (brt *BlockDBReadTransaction) GetJustifiedState() (state.State, error) {
	stateBytes, err := getKey(brt, justifiedStateKey)
	if err != nil {
		return nil, err
	}
	s := state.NewEmptyState()
	err = s.Unmarshal(stateBytes)
	return s, err
}

var blockRowPrefix = []byte("block-row")

// SetBlockRow sets a block row on disk to store the block chainindex.
func (but *BlockDBUpdateTransaction) SetBlockRow(disk *BlockNodeDisk) error {
	key := append(blockRowPrefix, disk.Hash[:]...)
	diskSer, err := disk.Marshal()
	if err != nil {
		return err
	}
	return setKey(but, key, diskSer)
}

// GetBlockRow gets the block row on disk.
func (brt *BlockDBReadTransaction) GetBlockRow(c chainhash.Hash) (*BlockNodeDisk, error) {
	key := append(blockRowPrefix, c[:]...)
	diskSer, err := getKey(brt, key)
	if err != nil {
		return nil, err
	}
	d := new(BlockNodeDisk)
	err = d.Unmarshal(diskSer)
	return d, err
}

var justifiedHeadKey = []byte("justified-head")

// SetJustifiedHead sets the latest justified head.
func (but *BlockDBUpdateTransaction) SetJustifiedHead(c chainhash.Hash) error {
	return setKeyHash(but, justifiedHeadKey, c)
}

// GetJustifiedHead gets the latest justified head.
func (brt *BlockDBReadTransaction) GetJustifiedHead() (chainhash.Hash, error) {
	return getKeyHash(brt, justifiedHeadKey)
}

var finalizedHeadKey = []byte("finalized-head")

// SetFinalizedHead sets the finalized head of the blockchain.
func (but *BlockDBUpdateTransaction) SetFinalizedHead(c chainhash.Hash) error {
	return setKeyHash(but, finalizedHeadKey, c)
}

// GetFinalizedHead gets the finalized head of the blockchain.
func (brt *BlockDBReadTransaction) GetFinalizedHead() (chainhash.Hash, error) {
	return getKeyHash(brt, finalizedHeadKey)
}

var genesisTimeKey = []byte("genesisTime")

// SetGenesisTime sets the genesis time of the blockchain.
func (but *BlockDBUpdateTransaction) SetGenesisTime(t time.Time) error {
	bs, err := t.MarshalBinary()
	if err != nil {
		return err
	}
	return setKey(but, genesisTimeKey, bs)
}

// GetGenesisTime gets the genesis time of the blockchain.
func (brt *BlockDBReadTransaction) GetGenesisTime() (time.Time, error) {
	bs, err := getKey(brt, genesisTimeKey)
	if err != nil {
		return time.Time{}, err
	}

	var t time.Time
	err = t.UnmarshalBinary(bs)
	return t, err
}

var _ DB = &blockDB{}
var _ DBUpdateTransaction = &BlockDBUpdateTransaction{}
var _ DBViewTransaction = &BlockDBReadTransaction{}

// DB is a database for storing chain state.
type DB interface {
	Close()
	Update(func(DBUpdateTransaction) error) error
	View(func(DBViewTransaction) error) error
}

// DBViewTransaction is a transaction to view the state of the database.
type DBViewTransaction interface {
	GetBlock(hash chainhash.Hash) (*primitives.Block, error)
	GetRawBlock(hash chainhash.Hash) ([]byte, error)
	GetTip() (chainhash.Hash, error)
	GetFinalizedState() (state.State, error)
	GetJustifiedState() (state.State, error)
	GetBlockRow(chainhash.Hash) (*BlockNodeDisk, error)
	GetJustifiedHead() (chainhash.Hash, error)
	GetFinalizedHead() (chainhash.Hash, error)
	GetGenesisTime() (time.Time, error)
}

// DBUpdateTransaction is a transaction to update the state of the database.
type DBUpdateTransaction interface {
	AddRawBlock(block *primitives.Block) error
	SetTip(chainhash.Hash) error
	SetFinalizedState(state.State) error
	SetJustifiedState(state.State) error
	SetBlockRow(*BlockNodeDisk) error
	SetJustifiedHead(chainhash.Hash) error
	SetFinalizedHead(chainhash.Hash) error
	SetGenesisTime(time.Time) error
	DBViewTransaction
}
