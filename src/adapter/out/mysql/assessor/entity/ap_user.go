package entity

import "database/sql"

type ApUserMySql struct {
	Id          int64          `json:"id"`
	UserCode    sql.NullString `json:"user_code"`
	UserName    sql.NullString `json:"user_name"`
	UserType    string         `json:"user_type"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Photo       string         `json:"photo"`
	PhoneNumber string         `json:"phone_number"`
}
