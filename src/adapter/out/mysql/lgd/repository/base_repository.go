package repository

import (
	"tandur.com/src/adapter/out/mysql/lgd/entity"
	"tandur.com/src/domain"
)

type LgdBaseRepository interface {
	GetAll(filter domain.Lgd) (*[]entity.LgdMySql, error)
	GetById(id int64) (*entity.LgdMySql, error)
	GetByToken(token string) (*[]entity.LgdMySql, error)
}
