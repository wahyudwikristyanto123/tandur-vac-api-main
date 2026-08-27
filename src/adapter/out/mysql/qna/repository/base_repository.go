package repository

import (
	"tandur.com/src/adapter/out/mysql/qna/entity"
	"tandur.com/src/domain"
)

type QnaBaseRepository interface {
	GetAll(filter domain.Qna) (*[]entity.QnaMySql, error)
	GetById(id int64) (*entity.QnaMySql, error)
	GetDetailsById(id int64) (*[]entity.QnaDetailMySql, error)
	GetByToken(token string) (*[]entity.QnaMySql, error)
	UpsertResult(data entity.QnaResultRequestMysql) error
	GetResultsByToken(token string) (*[]entity.QnaResultRequestMysql, error)
}
