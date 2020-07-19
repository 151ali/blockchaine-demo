package main

// Blockchain :
type Blockchain struct {
	Blocks []Block `json:"blockchaine"`
}

// NewBlockchain :
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []Block{*GenesisBlock()},
	}
}

// GetBlocks :
func (bc *Blockchain) GetBlocks() []Block {
	return bc.Blocks
}

// GetLastBlock :
func (bc *Blockchain) GetLastBlock() *Block {
	return &bc.Blocks[len(bc.Blocks)-1]
}

// GenesisBlock :
func GenesisBlock() *Block {
	return NewBlock(1806, "Genesis-Block", "")
}

// AddBlock :
func (bc *Blockchain) AddBlock(data string) {

	lastBlock := bc.GetLastBlock()
	newBlock := NewBlock(1, data, lastBlock.GetHash())

	bc.Blocks = append(bc.Blocks, *newBlock)
}

// VerifyBlockchaine :
func (bc *Blockchain) VerifyBlockchaine() error {
	// TODO :
	return nil
}
