package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wesley-lewis/go-blockchain/crypto"
	"github.com/wesley-lewis/go-blockchain/types"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Height:    height,
		Timestamp: time.Now().UnixNano(),
	}

	tx := Transaction{
		Data: []byte("hello"),
	}

	return NewBlock(header, []Transaction{tx})
}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height)

	assert.Nil(t, b.Sign(privKey))
	return b
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0)

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}
