package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
		"insert into %s (name, email, password_hash, blocked, created_at) values ($1, $2, $3, $4, $5) returning id", userTable,
	)
	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Blocked, user.CreatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(name, password string) (gym_app.User, error) {
	var user gym_app.User

	query := fmt.Sprintf("select id from %s where name=$1 and password_hash=$2", userTable)
	// TODO: refactor user to username
	err := r.db.Get(&user, query, name, password)
	logrus.Print(query)
	logrus.Print(err)
	return user, err
}
