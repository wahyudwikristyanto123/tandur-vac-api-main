package adapter

import "tandur.com/src/domain"

type RoleplayBaseAdapter interface {
	GetAll(filter domain.RolePlay) (*[]domain.RolePlay, error)
	GetById(id int64) (*domain.RolePlay, error)
	GetByToken(token string) (*[]domain.RolePlay, error)
}
