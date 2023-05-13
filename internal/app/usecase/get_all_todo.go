package usecase

import (
	"github.com/wellingtonlope/todo-heroku-api/internal/app/repository"
)

type GetAllTodo interface {
	Handle() ([]TodoOutput, error)
}

func NewGetAllTodo(todoRepository repository.Todo) GetAllTodo {
	return &getAllTodo{repository: todoRepository}
}

type getAllTodo struct {
	repository repository.Todo
}

func (uc *getAllTodo) Handle() ([]TodoOutput, error) {
	todos, err := uc.repository.GetAll()
	if err != nil {
		return []TodoOutput{}, err
	}

	return todoOutputsFromTodos(todos), nil
}
