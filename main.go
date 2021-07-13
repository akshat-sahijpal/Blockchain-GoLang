package main

import (
	"digicoin/blocks"
	"fmt"
	"strconv"
)

func main() {
	blockChain := blocks.InitializeBlockChain()
	blockChain.AddBlock("As")
	blockChain.AddBlock("As3")
	blockChain.AddBlock("As34")
	blockChain.AddBlock("As2")
	blockChain.AddBlock("A5s")
	for _, u := range blockChain.Blocks {
		fmt.Printf("\nHash Of Current Block = %x\nData Of Current Block = %s\nPrevious Hash Of Current Block = %x\n\n\n", u.Hash, u.Data, u.PrevHash)
		pow := blocks.POW(u)
		fmt.Printf("\nPOW = %s\n", strconv.FormatBool(pow.Validate()))
	}
}
