package blocks

import (
	"bytes"
	"crypto/sha256"
)

// Block Structure
type Block struct {
	Hash     []byte `json:"hash"`      // Hash of current block
	PrevHash []byte `json:"prev_hash"` // Hash of prev block
	Data     []byte `json:"data"`      // data of current block
}

// GenerateHash Hash of current block = PrevHash + Hash(Data)
func (bloc *Block) GenerateHash() {
	// Join concatenates PrevHash and Data together
	var joined []byte = bytes.Join([][]byte{(*bloc).PrevHash, (*bloc).Data}, []byte{})
	hash := sha256.Sum256(joined)
	(*bloc).Hash = hash[:]
}

// Instantiates a single block
func initBlock(blockData, prevHash []byte) *Block {
	var bloc *Block = &Block{
		Hash:     []byte{},
		PrevHash: prevHash,
		Data:     blockData,
	}
	bloc.GenerateHash() // Appends Hash Value to Current Block
	return bloc
}

// BlockChain

type BlockChain struct { // BlockChain is Basically Blocks Linked Together
	Blocks []*Block // Array of block pointers
}

func (blockChain *BlockChain) AddBlock(data string) {
	var prevBlock *Block = blockChain.Blocks[len(blockChain.Blocks)-1] // Previous Block in BChain
	var currentBlock *Block = initBlock([]byte(data), prevBlock.Hash)  // Creates A New Block
	blockChain.Blocks = append(blockChain.Blocks, currentBlock)
}

// GenesisBlock This block is the first block in the chain
func GenesisBlock() *Block {
	return initBlock([]byte("GenesisBlock"), []byte{})
}

// InitializeBlockChain Starts Blockchain
func InitializeBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{GenesisBlock()}}
}
