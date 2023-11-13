package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wesley-lewis/go-blockchain/types"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    2,
		Nonce:     150135,
	}

	buf := &bytes.Buffer{}

	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}

	assert.Nil(t, hDecode.DecodeBinary(buf))

	assert.Equal(t, h, hDecode)
}

func TestBlock(t *testing.T) {
	h := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    2,
		Nonce:     150135,
	}

	block := &Block{
		Header:      h,
		Transaction: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, block.EncodeBinary(buf))

	b := &Block{}
	assert.Nil(t, b.DecodeBinary(buf))

	assert.Equal(t, block, b)

	fmt.Printf("%+v", b)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			Timestamp: time.Now().UnixNano(),
			Height:    2,
			Nonce:     150135,
		},
		Transaction: []Transaction{},
	}

	h := b.Hash()
	fmt.Println(h)
	assert.False(t, h.IsZero())
}
