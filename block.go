package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block :
type Block struct {
	Proof     int    `json:"proof"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	PrevHash  string `json:"prevHash"`
	Hash      string `json:"hash"`
}

// NewBlock :
func NewBlock(p int, data string, prevBlockHash string) *Block {
	block := &Block{
		Proof:     p,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevBlockHash,
		Hash:      "",
	}
	block.SetHash()

	return block
}

// SetHash :
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	v := string(timestamp) + b.PrevHash + b.Data

	hash := sha256.Sum256([]byte(v))
	b.Hash = fmt.Sprintf("%x", hash[:])
}

// GetProof :
func (b *Block) GetProof() int {
	return b.Proof
}

// GetHash :
func (b *Block) GetHash() string {
	return b.Hash
}

// GetPrevHash :
func (b *Block) GetPrevHash() string {
	return b.PrevHash
}

// GetData :
func (b *Block) GetData() string {
	return b.Data
}

// ShowBlock :
func (b *Block) ShowBlock() {
	fmt.Println("--------------------")
	fmt.Println("Proof      : ", b.Proof)
	fmt.Println("Created at : ", time.Unix(b.Timestamp, 0))
	fmt.Println("Data       : ", b.Data)
	fmt.Printf("Hash       : %s\n", b.Hash[:5])
	fmt.Println("--------------------")
}
