package usecase

import "tandur.com/src/domain"

type OneOnOneUseCase interface {
	GetByToken(token string) (*[]domain.OneOnOne, error)
}
