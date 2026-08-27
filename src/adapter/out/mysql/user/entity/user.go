package entity

import "database/sql"

type UserMySql struct {
	Token    sql.NullString `json:"token"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Name     string         `json:"name"`
	Asesi    string         `json:"asesi"`
}
