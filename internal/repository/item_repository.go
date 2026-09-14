package repository

import (
	"sync"

	"example.com/m/v2/internal/domain"
)

type ItemRepository struct {
	mu    sync.RWMutex
	items map[int]domain.Item
}

func NewRepository() domain.ItemRepository {
	return &ItemRepository{
		items: make(map[int]domain.Item),
	}
}

func (r *ItemRepository) Create(item *domain.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.items[item.ID]; exists {
		return domain.ErrAlreadyExist
	}
	r.items[item.ID] = *item
	return nil
}

func (r *ItemRepository) FindByID(id int) (*domain.Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, exists := r.items[id]
	if !exists {
		return nil, domain.ErrNotFound
	}
	return &item, nil
}
