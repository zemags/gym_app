package repository

import (
	"github.com/jmoiron/sqlx"
	gym_app "github.com/zemags/gym_app/store"
)

type Authorization interface {
	CreateUser(user gym_app.User) (int, error)
	GetUser(user, password string) (gym_app.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
