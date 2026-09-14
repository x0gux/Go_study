package handler

import (
	"net/http"
	"strconv"

	"example.com/m/v2/internal/domain"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	itemService domain.ItemService
}

func NewItemHandler(i domain.ItemService) *ItemHandler {
	return &ItemHandler{
		itemService: i,
	}
}

func (i *ItemHandler) CreateItem(c echo.Context) error {
	var item domain.Item
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	if err := i.itemService.CreateItem(&item); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create item"})
	}
	return c.JSON(http.StatusCreated, item)
}

func (i *ItemHandler) FoundItemById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid item ID"})
	}
	item, err := i.itemService.GetItemByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Item not found"})
	}
	return c.JSON(http.StatusOK, item)
}
