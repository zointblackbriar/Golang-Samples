package main 

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	_ "time"
)


type Block struct {
	Timestamp int64
	Data []byte 
	PreviousHashBlock []byte
	Hash []byte 
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]bytes{b.PreviousHashBlock, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}



