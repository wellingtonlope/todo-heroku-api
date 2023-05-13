package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/usecase"
	http2 "github.com/wellingtonlope/todo-heroku-api/internal/infra/http"
	"github.com/wellingtonlope/todo-heroku-api/internal/infra/memory"
	"log"
	"os"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil && shouldReadEnvFile() {
		log.Fatalf("Error loading .env file: %v", err)
	}

	memoryRepository := memory.NewTodo()
	createTodo := usecase.NewCreateTodo(memoryRepository)
	getAllTodo := usecase.NewGetAllTodo(memoryRepository)
	controller := http2.NewTodo(createTodo, getAllTodo)

	e.POST("/todos", controller.Create)
	e.GET("/todos", controller.GetAll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func shouldReadEnvFile() bool {
	return os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == ""
}
