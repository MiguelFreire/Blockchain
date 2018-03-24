package main

import (
	"fmt"
	"time"
)

type Blockchain struct {
	nodeID string
	blocks []*Block
}

func newBlockChain(nodeID string) *Blockchain {
	blockchain := new(Blockchain)
	blockchain.nodeID = nodeID
	blockchain.blocks = append(blockchain.blocks, generateGenesisBlock())
	return blockchain
}

func (bc *Blockchain) addBlock(data string) {

	bcSize := len(bc.blocks) - 1

	prevBlock := bc.blocks[bcSize]

	newBlock := newBlock(prevBlock, data)

	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) print() {
	for _, b := range bc.blocks {
		fmt.Printf("Block number: %d\n", b.blockNum)
		fmt.Printf("Created at: %s\n", b.createdAt.Format(time.RFC3339))
		fmt.Printf("Data: %s\n", b.data)
		fmt.Printf("Hash: %x\n", b.hash)
		fmt.Printf("Previous Hash: %x\n", b.prevBlockHash)
		fmt.Println()
	}
}
