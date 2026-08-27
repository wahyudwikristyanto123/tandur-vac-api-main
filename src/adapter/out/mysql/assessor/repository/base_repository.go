package repository

import (
	"tandur.com/src/adapter/out/mysql/assessor/entity"
	"tandur.com/src/domain"
)

type AssessorBaseRepository interface {
	GetAll(filter domain.Assessor) (*[]entity.ApUserMySql, error)
	GetById(id int64) (*entity.ApUserMySql, error)
}
