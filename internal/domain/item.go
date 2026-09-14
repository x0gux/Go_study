package domain

type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ItemRequest struct {
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
	Create(item *Item) error
	FindByID(id int) (*Item, error)
}

type ItemService interface {
	CreateItem(item *Item) error
	GetItemByID(id int) (*Item, error)
}
