// Code generated by fastssz. DO NOT EDIT.
package primitives

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlockHeader object
func (b *BlockHeader) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlockHeader object to a target array
func (b *BlockHeader) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Version'
	dst = ssz.MarshalUint64(dst, b.Version)

	// Field (1) 'Nonce'
	dst = ssz.MarshalUint64(dst, b.Nonce)

	// Field (2) 'TxMerkleRoot'
	dst = append(dst, b.TxMerkleRoot[:]...)

	// Field (3) 'VoteMerkleRoot'
	dst = append(dst, b.VoteMerkleRoot[:]...)

	// Field (4) 'DepositMerkleRoot'
	dst = append(dst, b.DepositMerkleRoot[:]...)

	// Field (5) 'ExitMerkleRoot'
	dst = append(dst, b.ExitMerkleRoot[:]...)

	// Field (6) 'VoteSlashingMerkleRoot'
	dst = append(dst, b.VoteSlashingMerkleRoot[:]...)

	// Field (7) 'RANDAOSlashingMerkleRoot'
	dst = append(dst, b.RANDAOSlashingMerkleRoot[:]...)

	// Field (8) 'ProposerSlashingMerkleRoot'
	dst = append(dst, b.ProposerSlashingMerkleRoot[:]...)

	// Field (9) 'GovernanceVotesMerkleRoot'
	dst = append(dst, b.GovernanceVotesMerkleRoot[:]...)

	// Field (10) 'PrevBlockHash'
	dst = append(dst, b.PrevBlockHash[:]...)

	// Field (11) 'Timestamp'
	dst = ssz.MarshalUint64(dst, b.Timestamp)

	// Field (12) 'Slot'
	dst = ssz.MarshalUint64(dst, b.Slot)

	// Field (13) 'StateRoot'
	dst = append(dst, b.StateRoot[:]...)

	// Field (14) 'FeeAddress'
	dst = append(dst, b.FeeAddress[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BlockHeader object
func (b *BlockHeader) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 372 {
		return ssz.ErrSize
	}

	// Field (0) 'Version'
	b.Version = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'Nonce'
	b.Nonce = ssz.UnmarshallUint64(buf[8:16])

	// Field (2) 'TxMerkleRoot'
	copy(b.TxMerkleRoot[:], buf[16:48])

	// Field (3) 'VoteMerkleRoot'
	copy(b.VoteMerkleRoot[:], buf[48:80])

	// Field (4) 'DepositMerkleRoot'
	copy(b.DepositMerkleRoot[:], buf[80:112])

	// Field (5) 'ExitMerkleRoot'
	copy(b.ExitMerkleRoot[:], buf[112:144])

	// Field (6) 'VoteSlashingMerkleRoot'
	copy(b.VoteSlashingMerkleRoot[:], buf[144:176])

	// Field (7) 'RANDAOSlashingMerkleRoot'
	copy(b.RANDAOSlashingMerkleRoot[:], buf[176:208])

	// Field (8) 'ProposerSlashingMerkleRoot'
	copy(b.ProposerSlashingMerkleRoot[:], buf[208:240])

	// Field (9) 'GovernanceVotesMerkleRoot'
	copy(b.GovernanceVotesMerkleRoot[:], buf[240:272])

	// Field (10) 'PrevBlockHash'
	copy(b.PrevBlockHash[:], buf[272:304])

	// Field (11) 'Timestamp'
	b.Timestamp = ssz.UnmarshallUint64(buf[304:312])

	// Field (12) 'Slot'
	b.Slot = ssz.UnmarshallUint64(buf[312:320])

	// Field (13) 'StateRoot'
	copy(b.StateRoot[:], buf[320:352])

	// Field (14) 'FeeAddress'
	copy(b.FeeAddress[:], buf[352:372])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlockHeader object
func (b *BlockHeader) SizeSSZ() (size int) {
	size = 372
	return
}

// HashTreeRoot ssz hashes the BlockHeader object
func (b *BlockHeader) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlockHeader object with a hasher
func (b *BlockHeader) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Version'
	hh.PutUint64(b.Version)

	// Field (1) 'Nonce'
	hh.PutUint64(b.Nonce)

	// Field (2) 'TxMerkleRoot'
	hh.PutBytes(b.TxMerkleRoot[:])

	// Field (3) 'VoteMerkleRoot'
	hh.PutBytes(b.VoteMerkleRoot[:])

	// Field (4) 'DepositMerkleRoot'
	hh.PutBytes(b.DepositMerkleRoot[:])

	// Field (5) 'ExitMerkleRoot'
	hh.PutBytes(b.ExitMerkleRoot[:])

	// Field (6) 'VoteSlashingMerkleRoot'
	hh.PutBytes(b.VoteSlashingMerkleRoot[:])

	// Field (7) 'RANDAOSlashingMerkleRoot'
	hh.PutBytes(b.RANDAOSlashingMerkleRoot[:])

	// Field (8) 'ProposerSlashingMerkleRoot'
	hh.PutBytes(b.ProposerSlashingMerkleRoot[:])

	// Field (9) 'GovernanceVotesMerkleRoot'
	hh.PutBytes(b.GovernanceVotesMerkleRoot[:])

	// Field (10) 'PrevBlockHash'
	hh.PutBytes(b.PrevBlockHash[:])

	// Field (11) 'Timestamp'
	hh.PutUint64(b.Timestamp)

	// Field (12) 'Slot'
	hh.PutUint64(b.Slot)

	// Field (13) 'StateRoot'
	hh.PutBytes(b.StateRoot[:])

	// Field (14) 'FeeAddress'
	hh.PutBytes(b.FeeAddress[:])

	hh.Merkleize(indx)
	return
}
