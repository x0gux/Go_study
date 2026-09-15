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

// CreateItem godoc
// @Summary      아이템 생성
// @Description  새로운 아이템을 등록합니다.
// @Tags         items
// @Accept       json
// @Produce      json
// @Param        item  body      domain.ItemRequest  true  "생성할 아이템 정보"
// @Success      201   {object}  domain.Item
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /items [post]
func (i *ItemHandler) CreateItem(c echo.Context) error {
	var item domain.ItemRequest
	if err := c.Bind(&item); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}
	createdItem, err := i.itemService.CreateItem(&item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create item"})
	}

	returnItem := domain.ItemResponse{
		ID:          createdItem.ID,
		Name:        createdItem.Name,
		Description: createdItem.Description,
		Price:       createdItem.Price,
	}

	return c.JSON(http.StatusCreated, returnItem)
}

// FoundItemById godoc
// @Summary      아이템 단건 조회
// @Description  ID로 특정 아이템을 조회합니다.
// @Tags         items
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  domain.Item
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /items/{id} [get]
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

// GetAllItem godoc
// @Summary      아이템 전체 조회
// @Description  모든 아이템을 조회합니다.
// @Tags         items
// @Produce      json
// @Success      200  {object}  []domain.Item
// @Failure      500  {object}  map[string]interface{}
// @Router       /items [get]
func (i *ItemHandler) GetAllItem(c echo.Context) error {
	item, err := i.itemService.GetAllItem()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to get all items"})
	}
	return c.JSON(http.StatusOK, item)
}

// DeleteByID godoc
// @Summary      아이템 삭제
// @Description  ID로 특정 아이템을 삭제합니다.
// @Tags         items
// @Produce      json
// @Param        id   path      int  true  "Item ID"
// @Success      200  {object}  domain.Item
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /items/{id} [delete]
func (i *ItemHandler) DeleteByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid item ID"})
	}
	item, err := i.itemService.DeleteByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Item not found"})
	}
	return c.JSON(http.StatusOK, item)
}
