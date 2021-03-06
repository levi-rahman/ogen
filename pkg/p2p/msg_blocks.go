package p2p

import (
	"github.com/olympus-protocol/ogen/pkg/primitives"
)

// MaxBlocksPerMsg defines the maximum amount of blocks that a peer can send.
const MaxBlocksPerMsg = 32

// MsgBlocks is the struct of the message the is transmitted upon the network.
type MsgBlocks struct {
	Blocks []*primitives.Block `ssz-max:"32"`
}

// Marshal serializes the data to bytes
func (m *MsgBlocks) Marshal() ([]byte, error) {
	return m.MarshalSSZ()
}

// Unmarshal deserializes the data
func (m *MsgBlocks) Unmarshal(b []byte) error {
	return m.UnmarshalSSZ(b)
}

// Command returns the message topic
func (m *MsgBlocks) Command() string {
	return MsgBlocksCmd
}

// MaxPayloadLength returns the maximum size of the MsgBlocks message.
func (m *MsgBlocks) MaxPayloadLength() uint64 {
	return primitives.MaxBlockSize * MaxBlocksPerMsg
}
