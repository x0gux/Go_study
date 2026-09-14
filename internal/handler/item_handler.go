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
// @Param        item  body      domain.Item  true  "생성할 아이템 정보"
// @Success      201   {object}  domain.Item
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /items [post]
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
