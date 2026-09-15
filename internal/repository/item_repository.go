package repository

import (
	"example.com/m/v2/internal/domain"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.ItemRepository {
	db.AutoMigrate(&domain.Item{})
	return &ItemRepository{
		db: db,
	}

}

func (r *ItemRepository) Create(item *domain.ItemRequest) (*domain.Item, error) {
	r.db.Debug().Table("items").Create(item)
	return &domain.Item{}, nil
}

func (r *ItemRepository) GetAllItem() (*[]domain.Item, error) {
	var items []domain.Item
	r.db.Debug().Table("items").Find(&items)
	return &items, nil
}

func (r *ItemRepository) FindByID(id int) (*domain.Item, error) {
	var item domain.Item
	r.db.Debug().Where("id = ?", id).First(&item)
	if item.ID == 0 {
		return nil, domain.ErrNotFound
	}
	return &item, nil
}

func (r *ItemRepository) DeleteByID(id int) (*domain.Item, error) {
	var item domain.Item
	r.db.Debug().Where("id = ?", id).Delete(&item)
	if item.ID == 0 {
		return nil, domain.ErrNotFound
	}
	return &item, nil
}
