package repository

import (
	"tandur.com/src/adapter/out/mysql/subtes/entity"
	"tandur.com/src/domain"
)

type SubtesBaseRepository interface {
	GetAll(filter domain.Subtes) (*[]entity.SubtesMySql, error)
	GetById(id int64) (*entity.SubtesMySql, error)
	GetByToken(token string) (*[]entity.SubtesMySql, error)
	UpdateStatusById(id int64, status string) error
	UpdateResultById(id int64, result string) error
	GetResultById(id int64) (string, error)
}
