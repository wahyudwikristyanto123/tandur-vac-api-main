package usecase

import "tandur.com/src/domain"

type QnaUseCase interface {
	GetById(id int64) (*domain.Qna, error)
	GetByToken(token string) (*[]domain.Qna, error)
	UpsertResult(data domain.QnaResultRequest) error
	UpsertResults(data []domain.QnaResultRequest) error
	GetResultsByToken(token string) (*[]domain.QnaResultRequest, error)
}
