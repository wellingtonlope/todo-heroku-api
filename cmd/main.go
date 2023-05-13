package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/repository"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/usecase"
	http2 "github.com/wellingtonlope/todo-heroku-api/internal/infra/http"
	"github.com/wellingtonlope/todo-heroku-api/internal/infra/memory"
	"github.com/wellingtonlope/todo-heroku-api/internal/infra/postgres"
	"log"
	"os"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil && shouldReadEnvFile() {
		log.Fatalf("Error loading .env file: %v", err)
	}

	todoRepository := getRepository()
	createTodo := usecase.NewCreateTodo(todoRepository)
	getAllTodo := usecase.NewGetAllTodo(todoRepository)
	controller := http2.NewTodo(createTodo, getAllTodo)

	e.POST("/todos", controller.Create)
	e.GET("/todos", controller.GetAll)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func shouldReadEnvFile() bool {
	return os.Getenv("APP_ENV") == "local" || os.Getenv("APP_ENV") == ""
}

func getRepository() repository.Todo {
	source := os.Getenv("DATABASE_SOURCE")
	if source == "postgres" {
		postgresRepository, err := postgres.NewTodo(postgres.DatabaseConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DATABASE"),
		})
		if err != nil {
			log.Fatalf("Error loading postgres: %v", err)
		}
		return postgresRepository
	}

	return memory.NewTodo()
}
