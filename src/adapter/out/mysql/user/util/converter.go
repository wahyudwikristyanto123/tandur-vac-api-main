package util

import (
	"tandur.com/src/adapter/out/mysql/user/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(data entity.UserMySql) domain.User {
	return domain.User{
		Token:    data.Token.String,
		Username: data.Username,
		Email:    data.Email,
		Name:     data.Asesi,
	}
}
