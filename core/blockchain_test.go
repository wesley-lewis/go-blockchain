package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(1))
	assert.Nil(t, err)

	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	totalBlocks := 10
	for i := 2; i <= totalBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(totalBlocks))
	assert.Equal(t, len(bc.headers), totalBlocks)

	assert.NotNil(t, bc.AddBlock(randomBlock(20)))
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(1))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	lenBlocks := 1000

	for i := 2; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(uint32(i))
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}
}

func TestAddBlockTooHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3)))
}
