// Code generated by fastssz. DO NOT EDIT.
package p2p

import (
	ssz "github.com/ferranbt/fastssz"
	"github.com/olympus-protocol/ogen/pkg/primitives"
)

// MarshalSSZ ssz marshals the MsgExit object
func (m *MsgExit) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(m)
}

// MarshalSSZTo ssz marshals the MsgExit object to a target array
func (m *MsgExit) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Data'
	if m.Data == nil {
		m.Data = new(primitives.Exit)
	}
	if dst, err = m.Data.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the MsgExit object
func (m *MsgExit) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 192 {
		return ssz.ErrSize
	}

	// Field (0) 'Data'
	if m.Data == nil {
		m.Data = new(primitives.Exit)
	}
	if err = m.Data.UnmarshalSSZ(buf[0:192]); err != nil {
		return err
	}

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the MsgExit object
func (m *MsgExit) SizeSSZ() (size int) {
	size = 192
	return
}

// HashTreeRoot ssz hashes the MsgExit object
func (m *MsgExit) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(m)
}

// HashTreeRootWith ssz hashes the MsgExit object with a hasher
func (m *MsgExit) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'Data'
	if err = m.Data.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}