package repository

import (
	"github.com/wellingtonlope/todo-heroku-api/internal/domain"
)

type Todo interface {
	Insert(todo domain.Todo) (domain.Todo, error)
	GetAll() ([]domain.Todo, error)
}
