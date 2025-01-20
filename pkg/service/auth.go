package service

import (
	"crypto/sha1"
	"fmt"
	todoList "github.com/Ilhomiddin123"
	"github.com/Ilhomiddin123/pkg/repository"
	_ "hash"
)

const salt = "wertyuiojhg345678oiuyhgffy3"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todoList.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
