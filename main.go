package main

import (
	"digicoin/blocks"
	"fmt"
)

func main() {
	blockChain := blocks.InitializeBlockChain()
	blockChain.AddBlock("As")
	blockChain.AddBlock("As3")
	blockChain.AddBlock("As34")
	blockChain.AddBlock("As2")
	blockChain.AddBlock("A5s")
	for _, u := range blockChain.Blocks {
		fmt.Printf("\n%s, %s, %s\n", u.Hash, u.Data, u.PrevHash)
	}
}
