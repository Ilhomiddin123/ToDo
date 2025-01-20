package repository

import (
	todoList "github.com/Ilhomiddin123"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoList.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
