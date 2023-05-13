package memory

import (
	"github.com/google/uuid"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/repository"
	"github.com/wellingtonlope/todo-heroku-api/internal/domain"
)

type todo struct {
	todos []domain.Todo
}

func NewTodo() repository.Todo {
	return &todo{}
}

func (r *todo) Insert(todo domain.Todo) (domain.Todo, error) {
	todo.ID = uuid.New().String()

	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *todo) GetAll() ([]domain.Todo, error) {
	return r.todos, nil
}
