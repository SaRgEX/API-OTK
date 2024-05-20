package model

type User struct {
	Id         int    `json:"-" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Patronumic string `json:"patronumic" db:"patronumic"`
	Login      string `json:"login" db:"login"`
	Password   string `json:"password" db:"-"`
	Status     string `json:"status" db:"status"`
	Role       string `json:"role" db:"role"`
}

type UserOutput struct {
	Id         int    `json:"-" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Patronumic string `json:"patronumic" db:"patronumic"`
	Login      string `json:"login" db:"login"`
	Status     string `json:"status" db:"status"`
	Role       string `json:"role" db:"role"`
}
