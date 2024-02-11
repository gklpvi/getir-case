package storage

import (
	"getir-case/model"
	"sync"
)

// ------------------------------------------
// NOTE: This is a simple in-memory storage implementation
// we may use some external in-memory storage like Redis or Memcached
// but for the sake of simplicity and not "over engineering", we will use a simple in-memory storage with mutex lock
// ------------------------------------------

// InMemoryStorage implements Storage interface for in-memory storage
type InMemoryStorage struct {
	items map[string]model.DBRecord
	mu    sync.RWMutex
}

// NewInMemoryStorage creates a new InMemoryStorage instance
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		items: make(map[string]model.DBRecord),
	}
}

// GetAllItems retrieves all items from in-memory storage
func (s *InMemoryStorage) GetAllItems() ([]model.DBRecord, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var items []model.DBRecord
	for _, item := range s.items {
		items = append(items, item)
	}
	return items, nil
}

// AddItem adds a new item to in-memory storage
func (s *InMemoryStorage) AddItem(item model.DBRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	//
	//s.items[item.ID] = item
	//
	return nil
}
