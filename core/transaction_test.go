package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesley-lewis/go-blockchain/crypto"
)

func TestTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("hello"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}
