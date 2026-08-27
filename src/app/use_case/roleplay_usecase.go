package usecase

import "tandur.com/src/domain"

type RoleplayUseCase interface {
	GetByToken(token string) (*[]domain.RolePlay, error)
}
