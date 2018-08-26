package core

type BlockChain struct {
	 Blocks []*Block
}

// add new block to blockchain
func (bc *BlockChain) AddBlock(data string) {

	preBlock := bc.Blocks[len(bc.Blocks)-1]

	newBlock := NewBlock(data, preBlock.Hash)

	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}


