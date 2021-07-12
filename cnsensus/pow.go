package cnsensus

import (
	"bytes"
	"digicoin/blocks"
	"encoding/binary"
	"math/big"
)

/*
Proof Of Work is a Consensus Algorithm
Nonce + HashData = UniquePattern Of zeros (first few bytes only)

Nonce should be such that, the first few bytes of hash contains 0 (REQUIRED)
*/

const (
	difficulty = 10
	hashBytes  = 256
)

type proofOfWork struct {
	block        *blocks.Block // Block Reference
	targetOfWork *big.Int      // Target For Specifying Blockchain Conditions
}

func pow(block *blocks.Block) *proofOfWork {
	targ := big.NewInt(1) // Returns *Int of val 1 (targ is *Int) Number is large so its memory adr is stored
	/*
	   Left Shift Operation -> for eg   10 << 1 (Shifts the number by 1) becomes 20 [binary numbers digits are shifted by 1]
	   similarly,
	          func (z *Int) Lsh(x *Int, n uint) *Int
	              Lsh sets z = x << n and returns z.
	             targ = targ << 246
	   targ gets a very very big int
	*/
	targ = targ.Lsh(targ, uint(hashBytes-difficulty)) // Left Shift Operation

	return &proofOfWork{
		block:        block,
		targetOfWork: targ,
	}
}
func toHex(n int64) []byte {
	buffer := new(bytes.Buffer) // *Buffer: Will Store binary
	if err := binary.Write(buffer, binary.BigEndian, n); err != nil {
		panic("Error in :toHex")
	} // Writes n into buffer in binary representation
	return buffer.Bytes()
}
