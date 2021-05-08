package models

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Transaction struct {
	// from    []byte
	// to      []byte
	// amount  uint32
	timestamp int64
	message   string
	hash      string
}

func (t *Transaction) String() string {
	return t.hash + ", " + t.message
}

func NewTransaction(m string) *Transaction {
	t := Transaction{timestamp: time.Now().Unix(), message: m}
	t.hash = calculateTransactionHash(&t)

	return &t
}

func calculateTransactionHash(t *Transaction) string {
	timestamp := []byte(strconv.FormatInt(t.timestamp, 10))
	header := bytes.Join([][]byte{timestamp, []byte(t.message)}, []byte{})

	sum := sha256.Sum256(header)
	return "0x" + hex.EncodeToString(sum[:])
}
