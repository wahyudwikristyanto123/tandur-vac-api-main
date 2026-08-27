package repository

import (
	"tandur.com/src/adapter/out/mysql/roleplay/entity"
	"tandur.com/src/domain"
)

type RoleplayBaseRepository interface {
	GetAll(filter domain.RolePlay) (*[]entity.RoleplayMySql, error)
	GetById(id int64) (*entity.RoleplayMySql, error)
	GetByToken(token string) (*[]entity.RoleplayMySql, error)
}
