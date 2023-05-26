package namespace

import (
	"bytes"
	"encoding/hex"
)

type ID []byte

// Less returns true if nid < other, otherwise, false.
func (nid ID) Less(other ID) bool {
	return bytes.Compare(nid, other) < 0
}

// Equal returns true if nid == other, otherwise, false.
func (nid ID) Equal(other ID) bool {
	return bytes.Equal(nid, other)
}

// LessOrEqual returns true if nid <= other, otherwise, false.
func (nid ID) LessOrEqual(other ID) bool {
	return bytes.Compare(nid, other) <= 0
}

// Size returns the byte size of the nid.
func (nid ID) Size() IDSize {
	return IDSize(len(nid))
}

// String stringifies the nid.
func (nid ID) String() string {
	return string(nid)
}

// HexString returns hexadecimal encoding of nid.
func (nid ID) HexString() string {
	return hex.EncodeToString(nid)
}
