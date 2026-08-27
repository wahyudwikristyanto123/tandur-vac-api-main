package adapter

import "tandur.com/src/domain"

type OneOnOneBaseAdapter interface {
	GetAll(filter domain.OneOnOne) (*[]domain.OneOnOne, error)
	GetById(id int64) (*domain.OneOnOne, error)
	GetByToken(token string) (*[]domain.OneOnOne, error)
}
