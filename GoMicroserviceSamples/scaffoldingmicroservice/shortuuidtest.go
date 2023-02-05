package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/google/uuid"
	"github.com/lithammer/shortuuid/v4"
)

type base58Encoder struct {}

func (encoderSomething base58Encoder) Encode(u uuid.UUID) string {
	return base58.Encode(u[:])
}

func (encoderSomething base58Encoder) Decode(s string) (uuid.UUID, error) {
	return uuid.FromBytes(base58.Decode(s))
}

func main() {
	enc := base58Encoder{}
	fmt.Println("result is: ", shortuuid.NewWithEncoder(enc)) // sample encoding
}