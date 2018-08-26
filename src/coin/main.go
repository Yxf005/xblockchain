package main

import (
	"core"
	"fmt"
)

func main() {

	bc := core.NewBlockChain() // 初始化区块链
	bc.AddBlock("set 100 bitcoin to Alex")
	bc.AddBlock("set another 100 bitcoin to Alex ")

	for _, block := range bc.Blocks {
		fmt.Printf("pre hash: %x\n", block.PreBlockHash)
		fmt.Printf("Data: %s \n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}
