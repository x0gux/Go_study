package main

import (
	"log"

	"example.com/m/v2/internal/handler"
	"example.com/m/v2/internal/repository"
	"example.com/m/v2/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	itemRepo := repository.NewRepository()
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	port := "8080"

	e.POST("/items", itemHandler.CreateItem)
	e.GET("/items/:id", itemHandler.FoundItemById)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	err := e.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
