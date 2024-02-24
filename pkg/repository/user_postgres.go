package repository

import (
	"fmt"

	model "github.com/SaRgEX/Diplom/Model"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Find(id int) (model.UserOutput, error) {
	var user model.UserOutput
	query := fmt.Sprintf("SELECT id, first_name, last_name, patronumic, login, status, role FROM %s WHERE id = $1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}
