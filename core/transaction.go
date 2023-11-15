package core

import (
	"fmt"
	"io"

	"github.com/wesley-lewis/go-blockchain/crypto"
)

type Transaction struct {
	Data []byte

	From      crypto.PublicKey
	Signature *crypto.Signature
}

func (tx *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)

	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("Transaction has no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
