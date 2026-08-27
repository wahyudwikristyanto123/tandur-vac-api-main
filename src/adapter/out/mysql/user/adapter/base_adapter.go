package adapter

import "tandur.com/src/domain"

type UserBaseAdapter interface {
	GetAll(filter domain.User) (*[]domain.User, error)
	GetById(id int64) (*domain.User, error)
	GetByToken(token string) (*[]domain.User, error)
}
