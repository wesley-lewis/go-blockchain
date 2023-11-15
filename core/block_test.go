package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wesley-lewis/go-blockchain/crypto"
	"github.com/wesley-lewis/go-blockchain/types"
)

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, 0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 100
	assert.NotNil(t, b.Verify())
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(t, height, prevBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)

	assert.Nil(t, b.Sign(privKey))
	return b
}

func randomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	tx := randomTxWithSignature(t)
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     time.Now().UnixNano(),
	}
	b := NewBlock(header, []Transaction{*tx})
	assert.Nil(t, b.Sign(privKey))

	return b
}
