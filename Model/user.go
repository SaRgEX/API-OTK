package model

type Account struct {
	Id         int    `json:"-" db:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronumic string `json:"patronumic"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Status     string `json:"status"`
	Role       string `json:"role"`
}
