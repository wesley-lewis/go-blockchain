package core

import (
	"io"

	"github.com/wesley-lewis/go-blockchain/crypto"
)

type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}
