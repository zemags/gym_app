package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	gym_app "github.com/zemags/gym_app/store"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user gym_app.User) (int, error) {
	// user id
	var id int
	query := fmt.Sprintf(
		"insert into %s (username, email, password_hash, created_at) values ($1, $2, $3, $4) returning id", userTable,
	)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, user.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (gym_app.User, error) {
	var user gym_app.User

	query := fmt.Sprintf("select id from %s where username=$1 and password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
