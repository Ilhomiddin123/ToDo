package service

import (
	todoList "github.com/Ilhomiddin123"
	"github.com/Ilhomiddin123/pkg/repository"
)

type Authorization interface {
	CreateUser(user todoList.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
