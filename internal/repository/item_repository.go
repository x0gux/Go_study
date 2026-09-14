package repository

import (
	"sync"

	"example.com/m/v2/internal/domain"
)

type ItemRepository struct {
	mu     sync.RWMutex
	items  map[int]domain.Item
	nextID int
}

func NewRepository() domain.ItemRepository {
	return &ItemRepository{
		items: make(map[int]domain.Item),
	}
}

func (r *ItemRepository) Create(item *domain.ItemRequest) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.nextID++
	currentID := r.nextID
	if _, exists := r.items[currentID]; exists {
		return domain.ErrAlreadyExist
	}
	r.items[currentID] = domain.Item{
		ID:          currentID,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
	}
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

func (r *ItemRepository) DeleteByID(id int) (*domain.Item, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.items[id]; !exists {
		return nil, domain.ErrNotFound
	}
	item := r.items[id]
	delete(r.items, id)
	return &item, nil
}
