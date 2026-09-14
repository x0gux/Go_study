package service

import "example.com/m/v2/internal/domain"

type ItemService struct {
	itemRepository domain.ItemRepository
}

func NewItemService(i domain.ItemRepository) *ItemService {
	return &ItemService{
		itemRepository: i,
	}
}

func (s *ItemService) CreateItem(item *domain.Item) error {

	if item.Price <= 0 {
		return domain.ErrInvalid
	}

	return s.itemRepository.Create(item)
}

func (s *ItemService) GetItemByID(id int) (*domain.Item, error) {
	if id <= 0 {
		return nil, domain.ErrInvalid
	}
	return s.itemRepository.FindByID(id)
}
