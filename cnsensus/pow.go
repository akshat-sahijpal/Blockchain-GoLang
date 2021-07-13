package cnsensus

import (
	"bytes"
	"digicoin/blocks"
	"digicoin/util"
	"math/big"
)

/*
Proof Of Work is a Consensus Algorithm
Nonce + HashData = UniquePattern Of zeros (first few bytes only)

Nonce should be such that, the first few bytes of hash contains 0 (REQUIRED)
*/

const (
	difficulty = 10
	hashBytes  = 256 // SHA-256
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
func (prfWork *proofOfWork) genBlockData(nonce int) []byte {
	return bytes.Join([][]byte{
		prfWork.block.PrevHash,
		prfWork.block.Data,
		util.ToHex(int64(nonce)),
		util.ToHex(int64(difficulty)),
	}, []byte{})
}
