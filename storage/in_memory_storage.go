package storage

import (
	"getir-case/model"
	"log"
	"sync"
)

// ------------------------------------------
// NOTE: This is a simple in-memory storage implementation
// we may use some external in-memory storage like Redis or Memcached
// but for the sake of simplicity and not "over engineering", we will use a simple in-memory storage with mutex lock
// ------------------------------------------

// InMemoryStorage implements Storage interface for in-memory storage
type InMemoryStorage struct {
	items map[string]model.IMRecord
	mu    sync.RWMutex
}

// NewInMemoryStorage creates a new InMemoryStorage instance
func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		items: make(map[string]model.IMRecord),
	}
}

// GetAllItems retrieves all items from in-memory storage
func (s *InMemoryStorage) GetAllItems() []model.IMRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var items []model.IMRecord
	for _, item := range s.items {
		record := model.IMRecord{
			Key:   item.Key,
			Value: item.Value,
		}
		items = append(items, record)
	}
	return items
}

// GetItem retrieves an item from in-memory storage
func (s *InMemoryStorage) GetItem(key string) (model.IMRecord, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.items[key]
	if !ok {
		log.Println("item not found")
		return model.IMRecord{}, false
	}
	return item, true
}

// AddItem adds a new item to in-memory storage
func (s *InMemoryStorage) AddItem(item *model.IMRecord) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items[item.Key] = *item
}
