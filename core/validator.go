package core

import "fmt"

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contains block {%d} with hash {%s}", b.Height, b.Hash(BlockHasher{}))
	}

	if b.Height > v.bc.Height()+1 {
		return fmt.Errorf("block {%s} too high", b.Hash(BlockHasher{}))
	}

	_, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	if err := b.Verify(); err != nil {
		return err
	}
	return nil
}
