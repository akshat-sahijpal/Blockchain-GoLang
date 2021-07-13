package blocks

import (
	"bytes"
	"crypto/sha256"
	"digicoin/util"
	"math"
	"math/big"
)

/*
Proof Of Work is a Consensus Algorithm
Nonce + HashData = UniquePattern Of zeros (first few bytes only)

Nonce should be such that, the first few bytes of hash contains 0 (REQUIRED)
*/

const (
	difficulty = 20
	hashBytes  = 256 // SHA-256
)

type proofOfWork struct {
	block        *Block   // Block Reference
	targetOfWork *big.Int // Target For Specifying Blockchain Conditions
}

func POW(block *Block) *proofOfWork {
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

func (pow *proofOfWork) Compute() (int, []byte) {
	var intHash big.Int
	var Hash [32]byte

	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.genBlockData(nonce)
		Hash = sha256.Sum256(data)
		intHash.SetBytes(Hash[:])
		// Cmp: x < y -> -1, x == y -> 0, x > y -> +1
		if intHash.Cmp(pow.targetOfWork) == -1 { // intHash < target
			break
		} else {
			nonce++
		}
	}
	return nonce, Hash[:]
}

func (pow *proofOfWork) Validate() bool {
	var Hash big.Int
	dt := pow.genBlockData(pow.block.Nonce)
	hash := sha256.Sum256(dt)
	Hash.SetBytes(hash[:])
	return Hash.Cmp(pow.targetOfWork) == -1 // x < y
}

func (prfWork *proofOfWork) genBlockData(nonce int) []byte {
	return bytes.Join([][]byte{
		prfWork.block.PrevHash,
		prfWork.block.Data,
		util.ToHex(int64(nonce)),
		util.ToHex(int64(difficulty)),
	}, []byte{})
}
