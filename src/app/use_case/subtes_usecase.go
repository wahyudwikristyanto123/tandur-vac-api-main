package usecase

import "tandur.com/src/domain"

type SubtesUseCase interface {
	GetResultById(id int64) (string, error)
	GetByToken(token string) (*[]domain.Subtes, error)
	UpdateStatusById(id int64, status string) error
	SubmitResult(id int64, result string) error
}
