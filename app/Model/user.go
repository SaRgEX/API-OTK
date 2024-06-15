package model

type User struct {
	Id         int    `json:"-" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Patronumic string `json:"patronumic" db:"patronumic"`
	Login      string `json:"login" db:"login"`
	Password   string `json:"password" db:"-"`
}

type UserWithRole struct {
	User
	Role   string `json:"role" db:"role" binding:"required"`
	Status string `json:"status" db:"status" binding:"required"`
}

type UserOutput struct {
	Id         int    `json:"-" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Patronumic string `json:"patronumic" db:"patronumic"`
	Phone      string `json:"phone" db:"phone"`
	Email      string `json:"email" db:"email"`
	Login      string `json:"login" db:"login"`
	Status     string `json:"status" db:"status"`
	Role       string `json:"role" db:"role"`
}

type UserOutputAdmin struct {
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Patronumic string `json:"patronumic" db:"patronumic"`
	UserInfo
}

type UpdateUser struct {
	FirstName  *string `json:"first_name" db:"first_name"`
	LastName   *string `json:"last_name" db:"last_name"`
	Patronumic *string `json:"patronumic" db:"patronumic"`
	Phone      *string `json:"phone" db:"phone"`
	Email      *string `json:"email" db:"email"`
}

type UserPayload struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
}

type UserInfo struct {
	Phone string `json:"phone" db:"phone"`
	Email string `json:"email" db:"email"`
}
