package adapter

import (
	"tandur.com/src/domain"
)

type QnaBaseAdapter interface {
	GetAll(filter domain.Qna) (*[]domain.Qna, error)
	GetDetailsById(id int64) (*[]domain.QnaDetail, error)
	GetById(id int64) (*domain.Qna, error)
	GetByToken(token string) (*[]domain.Qna, error)
	UpsertResult(data domain.QnaResultRequest) error
	GetResultsByToken(token string) (*[]domain.QnaResultRequest, error)
}
