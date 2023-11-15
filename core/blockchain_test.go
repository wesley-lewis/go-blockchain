package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	totalBlocks := 1000
	for i := 0; i < 1000; i++ {
		block := randomBlockWithSignature(t, uint32(i+1))
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(totalBlocks))
	assert.Equal(t, len(bc.headers), totalBlocks+1)

	assert.NotNil(t, bc.AddBlock(randomBlock(20)))
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}
