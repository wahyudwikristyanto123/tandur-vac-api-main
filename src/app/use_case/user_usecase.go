package usecase

import "tandur.com/src/domain"

type UserUseCase interface {
	GetUserByToken(token string) (*[]domain.User, error)
}
