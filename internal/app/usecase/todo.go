package usecase

import "github.com/wellingtonlope/todo-heroku-api/internal/domain"

type TodoOutput struct {
	ID          string
	Title       string
	Description string
}

func todoOutputFromTodo(todo domain.Todo) TodoOutput {
	return TodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
}

func todoOutputsFromTodos(todos []domain.Todo) (outputs []TodoOutput) {
	for _, todo := range todos {
		outputs = append(outputs, todoOutputFromTodo(todo))
	}
	return
}
