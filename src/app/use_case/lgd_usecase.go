package usecase

import "tandur.com/src/domain"

type LgdUseCase interface {
	GetByToken(token string) (*[]domain.Lgd, error)
}
