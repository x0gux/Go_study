package main

import (
	"log"

	_ "example.com/m/v2/docs"
	"example.com/m/v2/internal/handler"
	"example.com/m/v2/internal/repository"
	"example.com/m/v2/internal/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := "host=localhost user=postgres password=your_password dbname=amr_db port=5432 sslmode=disable TimeZone=Asia/Seoul"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	itemRepo := repository.NewRepository(db)
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	port := "8080"

	e.POST("/items", itemHandler.CreateItem)
	e.GET("/items/:id", itemHandler.FoundItemById)
	e.DELETE("/items/:id", itemHandler.DeleteByID)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	errs := e.Start(":" + port)
	if errs != nil || err != nil {
		log.Fatal(errs, err)
	}
}
