package domain

import "errors"

var (
	ErrTodoTitleIsInvalid = errors.New("title mustn't be empty")
)

type Todo struct {
	ID          string
	Title       string
	Description string
}

func NewTodo(title, description string) (Todo, error) {
	if title == "" {
		return Todo{}, ErrTodoTitleIsInvalid
	}
	return Todo{
		Title:       title,
		Description: description,
	}, nil
}
