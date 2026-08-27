package repository

import (
	"tandur.com/src/adapter/out/mysql/one_on_one/entity"
	"tandur.com/src/domain"
)

type OneOnOneBaseRepository interface {
	GetAll(filter domain.OneOnOne) (*[]entity.OneOnOneMySql, error)
	GetById(id int64) (*entity.OneOnOneMySql, error)
	GetByToken(token string) (*[]entity.OneOnOneMySql, error)
}
