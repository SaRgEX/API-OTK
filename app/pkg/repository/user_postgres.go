package repository

import (
	"fmt"
	"strings"

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
	query := fmt.Sprintf("SELECT id, first_name, last_name, patronumic, phone, email, login, status, role FROM %s WHERE id = $1", userTable)
	err := r.db.Get(&user, query, id)
	return user, err
}

func (r *UserPostgres) UpdateUser(id int, input model.UpdateUser) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	ardId := 1

	if input.FirstName != nil {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", ardId))
		args = append(args, *input.FirstName)
		ardId++
	}

	if input.LastName != nil {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", ardId))
		args = append(args, *input.LastName)
		ardId++
	}

	if input.Patronumic != nil {
		setValues = append(setValues, fmt.Sprintf("patronumic=$%d", ardId))
		args = append(args, *input.Patronumic)
		ardId++
	}

	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", ardId))
		args = append(args, *input.Phone)
		ardId++
	}

	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", ardId))
		args = append(args, *input.Email)
		ardId++
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", userTable, strings.Join(setValues, ", "), ardId)

	args = append(args, id)

	_, err := r.db.Exec(query, args...)

	return err
}
