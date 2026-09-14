package main

import (
	"log"

	_ "example.com/m/v2/docs"
	"example.com/m/v2/internal/handler"
	"example.com/m/v2/internal/repository"
	"example.com/m/v2/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	e := echo.New()

	itemRepo := repository.NewRepository()
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	port := "8080"

	e.POST("/items", itemHandler.CreateItem)
	e.GET("/items/:id", itemHandler.FoundItemById)
	e.DELETE("/items/:id", itemHandler.DeleteByID)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	err := e.Start(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
