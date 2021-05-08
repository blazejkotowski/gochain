package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	timestamp    int64
	transactions []*Transaction
	hash         string
	previousHash string
}

func NewBlock(previousHash string, transactions []*Transaction) *Block {
	block := Block{
		timestamp:    time.Now().Unix(),
		transactions: transactions,
		previousHash: previousHash,
	}
	block.hash = calculateBlockHash(&block)

	return &block
}

func (b *Block) Hash() string {
	return b.hash
}

func (b *Block) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("\033[36mBlock %v\033[0m\n", b.hash))
	sb.WriteString(fmt.Sprintf("\ttimestamp:\t%v\n", b.timestamp))
	sb.WriteString(fmt.Sprintf("\tpreviousHash:\t%v\n", b.previousHash))

	sb.WriteString("\tTransactions: [\n")
	for _, t := range b.transactions {
		sb.WriteString("\t\t")
		sb.WriteString(t.String())
		sb.WriteString("\n")
	}
	sb.WriteString("\t]\n")

	return sb.String()
}

func calculateBlockHash(b *Block) string {
	transactionsHash := []byte(nil)

	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	header := bytes.Join([][]byte{timestamp, []byte(b.previousHash), transactionsHash}, []byte{})

	sum := sha256.Sum256(header)
	return "0x" + hex.EncodeToString(sum[:])
}
