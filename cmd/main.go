package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/usecase"
	http2 "github.com/wellingtonlope/todo-heroku-api/internal/infra/http"
	"github.com/wellingtonlope/todo-heroku-api/internal/infra/memory"
	"log"
	"os"
	"time"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil && shouldReadEnvFile() {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("todo-api-stage"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatalf("Error loading new relic file: %v", err)
	}
	app.RecordLog(newrelic.LogData{Message: "teste", Severity: "error"})
	app.RecordCustomEvent("just_a_test", map[string]interface{}{"name": "test"})
	trs := app.StartTransaction("a_transaction")
	time.Sleep(time.Second)
	trs.End()
	app.RecordCustomMetric("just_a_metric", 1.2)

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
