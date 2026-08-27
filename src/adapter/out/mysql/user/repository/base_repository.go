package repository

import (
	"tandur.com/src/adapter/out/mysql/user/entity"
	"tandur.com/src/domain"
)

type UserBaseRepository interface {
	GetAll(filter domain.User) (*[]entity.UserMySql, error)
	GetById(id int64) (*entity.UserMySql, error)
	GetByToken(token string) (*[]entity.UserMySql, error)
}
