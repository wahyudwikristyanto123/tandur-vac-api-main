package repository

import (
	"tandur.com/src/adapter/out/mysql/cv/entity"
)

type CvBaseRepository interface {
	GetCvByToken(token string) (*[]entity.CvViewMySql, error)
	GetCvById(id int64) (*entity.CvViewMySql, error)
	InsertCv(data entity.CvMySql) (int64, error)
	UpdateCv(data entity.CvMySql) error
	DeleteCv(id int64) error
}
