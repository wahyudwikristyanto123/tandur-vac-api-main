package usecase

import "tandur.com/src/domain"

type CvUseCase interface {
	GetByToken(token string) (*domain.Cv, error)
	SubmitCv(data domain.Cv) error
}
