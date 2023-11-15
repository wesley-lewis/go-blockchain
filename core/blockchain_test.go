package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesley-lewis/go-blockchain/types"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(t, 1, types.Hash{}))
	assert.Nil(t, err)

	return bc
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	totalBlocks := 1000
	for i := 2; i <= totalBlocks; i++ {
		// block := randomBlockWithSignature(t, uint32(i), getPrevBlockHash(t, bc, uint32(i)))
		block := randomBlockWithSignature(t, uint32(i), getPrevBlockHash(t, bc, uint32(i-1)))
		fmt.Println(block.Transactions[0].Signature)
		assert.Nil(t, bc.AddBlock(block))
	}

	assert.Equal(t, bc.Height(), uint32(totalBlocks))
	assert.Equal(t, len(bc.headers), totalBlocks)

	assert.NotNil(t, bc.AddBlock(randomBlock(t, 20, types.Hash{})))
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

	fmt.Println("TestGetHeader:", bc.Height())
	for i := 2; i < lenBlocks; i++ {
		block := randomBlockWithSignature(t, uint32(i), getPrevBlockHash(t, bc, uint32(i-1)))
		assert.Nil(t, bc.AddBlock(block))
		header, err := bc.GetHeader(uint32(i))
		assert.Nil(t, err)
		assert.Equal(t, header, block.Header)
	}
}

// TODO: Too high blocks can't be hashed thus produce an error while getting its prevHash

// func TestAddBlockTooHigh(t *testing.T) {
// 	bc := newBlockchainWithGenesis(t)
//
// 	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3, getPrevBlockHash(t, bc, uint32(2)))))
// }

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(prevHeader)
}
