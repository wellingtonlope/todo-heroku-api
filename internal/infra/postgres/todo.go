package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/wellingtonlope/todo-heroku-api/internal/app/repository"
	"github.com/wellingtonlope/todo-heroku-api/internal/domain"
)

type TodoDB struct {
	ID          string `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type todo struct {
	conn *sqlx.DB
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Schema   string
}

func NewTodo(dbConfig DatabaseConfig) (repository.Todo, error) {
	conn, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User,
		dbConfig.Password, dbConfig.DBName,
	))
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	return &todo{conn}, err
}

func (r *todo) Insert(todo domain.Todo) (domain.Todo, error) {
	todo.ID = uuid.NewString()
	tododb := TodoDB{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
	}
	_, err := r.conn.NamedExec("insert into public.todo (id, title, description) values (:id, :title, :description);", tododb)
	return todo, err
}

func (r *todo) GetAll() (todos []domain.Todo, err error) {
	err = r.conn.Select(&todos, "select id, title, description from public.todo")
	return
}
