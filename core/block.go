package core

import (
	"io"

	"github.com/wesley-lewis/go-blockchain/crypto"
	"github.com/wesley-lewis/go-blockchain/types"
)

type Header struct {
	Version   uint32
	DataHash  types.Hash
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32
}

type Block struct {
	Header
	Transaction []Transaction
	Validator   crypto.PublicKey
	Signature   *crypto.Signature

	// Cached version of the header hash
	hash types.Hash
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {

	return nil
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return nil
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}
