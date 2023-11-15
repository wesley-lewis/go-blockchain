package core

type Storage interface {
	Put(*Block) error
}

type MemoryStore struct {
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{}
}

// TODO
func (m *MemoryStore) Put(b *Block) error {
	return nil
}
