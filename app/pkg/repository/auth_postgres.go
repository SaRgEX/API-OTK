package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(account model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (first_name, last_name, patronumic, login, password_hash) values ($1, $2, $3, $4, $5) RETURNING id", userTable)
	row := r.db.QueryRow(query, account.FirstName, account.LastName, account.Patronumic, account.Login, account.Password)
	if account.Role != "" && account.Status != "" {
		query = fmt.Sprintf("INSERT INTO %s (first_name, last_name, patronumic, login, password_hash, status, role) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", userTable)
		row = r.db.QueryRow(query, account.FirstName, account.LastName, account.Patronumic, account.Login, account.Password, account.Status, account.Role)
	}
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(login, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, login, password)
	return user, err
}

func (r *AuthPostgres) Logout(token string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE token=$1", userTable)
	_, err := r.db.Exec(query, token)
	return err
}
