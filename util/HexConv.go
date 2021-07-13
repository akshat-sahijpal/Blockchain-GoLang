package util

import (
	"bytes"
	"encoding/binary"
)

func ToHex(n int64) []byte {
	buffer := new(bytes.Buffer) // *Buffer: Will Store binary
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		panic("Error in :toHex")
	} // Writes n into buffer in binary representation
	return buffer.Bytes()
}
