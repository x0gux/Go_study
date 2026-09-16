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
	itemreq := &domain.Item{
		Name:        item.Name,
		Price:       item.Price,
		Description: item.Description,
	}
	r.db.Debug().Table("items").Create(itemreq)
	if r.db.Error != nil {
		return nil, r.db.Error
	}
	return itemreq, nil
}

func (r *ItemRepository) GetAllItem() (*[]domain.Item, error) {
	var items []domain.Item
	r.db.Debug().Table("items").Order("ID ASC").Find(&items)
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

func (r *ItemRepository) ModifyById(id int, item *domain.ItemRequest) (*domain.Item, error) {
	var items domain.Item
	var modifyitem = r.db.Debug().Table("items").Where("id = ?", id).Updates(map[string]any{
		"name":        item.Name,
		"price":       item.Price,
		"description": item.Description,
	})

	if modifyitem.Error != nil {
		return nil, modifyitem.Error
	}
	return &items, nil
}
