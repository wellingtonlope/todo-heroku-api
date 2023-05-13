package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/usecase"
	"net/http"
)

type TodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoOutput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Todo interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
}

func NewTodo(createUC usecase.CreateTodo, getAllUC usecase.GetAllTodo) Todo {
	return &todo{createUC: createUC, getAllUC: getAllUC}
}

type todo struct {
	createUC usecase.CreateTodo
	getAllUC usecase.GetAllTodo
}

func (ctr *todo) Create(c echo.Context) error {
	todoInput := new(TodoInput)
	if err := c.Bind(todoInput); err != nil {
		return err
	}
	todoOutput, err := ctr.createUC.Handle(usecase.CreateTodoInput{
		Title:       todoInput.Title,
		Description: todoInput.Description,
	})
	if err != nil {
		log.Errorf("[func:GetAll] %v", err)
		return err
	}
	return c.JSON(http.StatusCreated, TodoOutput{
		ID:          todoOutput.ID,
		Title:       todoOutput.Title,
		Description: todoOutput.Description,
	})
}

func (ctr *todo) GetAll(c echo.Context) error {
	todoOutputs, err := ctr.getAllUC.Handle()
	if err != nil {
		log.Errorf("[func:GetAll] %v", err)
		return err
	}

	outputs := make([]TodoOutput, 0)
	for _, todoOutput := range todoOutputs {
		outputs = append(outputs, TodoOutput{
			ID:          todoOutput.ID,
			Title:       todoOutput.Title,
			Description: todoOutput.Description,
		})
	}

	return c.JSON(http.StatusOK, outputs)
}
