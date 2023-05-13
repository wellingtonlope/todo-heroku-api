package usecase

import (
	"github.com/wellingtonlope/todo-heroku-api/internal/app/repository"
	"github.com/wellingtonlope/todo-heroku-api/internal/domain"
)

type CreateTodoInput struct {
	Title       string
	Description string
}

type CreateTodo interface {
	Handle(input CreateTodoInput) (TodoOutput, error)
}

func NewCreateTodo(todoRepository repository.Todo) CreateTodo {
	return &createTodo{repository: todoRepository}
}

type createTodo struct {
	repository repository.Todo
}

func (uc *createTodo) Handle(input CreateTodoInput) (TodoOutput, error) {
	todo, err := domain.NewTodo(input.Title, input.Description)
	if err != nil {
		return TodoOutput{}, err
	}

	todo, err = uc.repository.Insert(todo)
	if err != nil {
		return TodoOutput{}, err
	}

	return todoOutputFromTodo(todo), nil
}
