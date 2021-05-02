package service

import (
	"github.com/zemags/gym_app/pkg/repository"
	gym_app "github.com/zemags/gym_app/store"
)

type Authorization interface {
	CreateUser(user gym_app.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
	}
}
