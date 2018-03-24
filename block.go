package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	blockNum      int
	createdAt     time.Time
	hash          string
	prevBlockHash string //previous block hash
	data          string
}

func newBlock(prevBlock *Block, data string) *Block {
	b := new(Block)
	b.blockNum = prevBlock.blockNum + 1
	b.createdAt = time.Now()
	b.data = data
	b.prevBlockHash = prevBlock.hash
	b.hash = b.hashBlock()

	return b
}

func generateGenesisBlock() *Block {
	b := new(Block)

	b.blockNum = 0
	b.createdAt = time.Now()
	b.data = ""
	b.prevBlockHash = ""
	b.hash = b.hashBlock()

	return b

}

func (b *Block) hashBlock() string {
	headers := string(b.blockNum) + b.createdAt.Format(time.UnixDate) + b.prevBlockHash + b.data
	hash := sha256.New()
	hash.Write([]byte(headers))

	final := hash.Sum(nil)

	return hex.EncodeToString(final)
}

func isBlockValid(new Block, prev Block) bool {

	if prev.blockNum+1 != new.blockNum {
		return false
	}

	if prev.hash != new.prevBlockHash {
		return false
	}

	if new.hash != new.hashBlock() {
		return false
	}

	return true
}
