package models

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	blocks   []*Block
	accounts map[string]*Account
}

func NewBlockchain() *Blockchain {
	blockchain := Blockchain{
		blocks:   []*Block{},
		accounts: make(map[string]*Account),
	}

	genesis := newGenesisBlock()
	blockchain.AddBlock(genesis)

	return &blockchain
}

func (bc *Blockchain) AddBlock(b *Block) {
	bc.blocks = append(bc.blocks, b)
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *Blockchain) String() string {
	var sb strings.Builder

	for i, b := range bc.blocks {
		sb.WriteString(fmt.Sprintf("Block %v\n", i+1))
		sb.WriteString(b.String())
		sb.WriteString("\n\n")
	}

	return sb.String()
}

func newGenesisBlock() *Block {
	return NewBlock("0x0000000000000000000000000000000000000000000000000000000000000000", nil)
}
