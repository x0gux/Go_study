package domain

type Item struct {
	ID          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ItemRequest struct {
	ID          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ItemResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ItemRepository interface {
	Create(item *ItemRequest) (*Item, error)
	FindByID(id int) (*Item, error)
	DeleteByID(id int) (*Item, error)
}

type ItemService interface {
	CreateItem(item *ItemRequest) (*Item, error)
	GetItemByID(id int) (*Item, error)
	DeleteByID(id int) (*Item, error)
}
