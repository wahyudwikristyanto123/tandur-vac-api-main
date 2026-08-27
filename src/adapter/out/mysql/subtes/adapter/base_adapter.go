package adapter

import "tandur.com/src/domain"

type SubtesBaseAdapter interface {
	GetAll(filter domain.Subtes) (*[]domain.Subtes, error)
	GetById(id int64) (*domain.Subtes, error)
	GetByToken(token string) (*[]domain.Subtes, error)
	UpdateStatusById(id int64, status string) error
	UpdateResultById(id int64, result string) error
	GetResultById(id int64) (string, error)
}
