package adapter

import "tandur.com/src/domain"

type LgdBaseAdapter interface {
	GetAll(filter domain.Lgd) (*[]domain.Lgd, error)
	GetById(id int64) (*domain.Lgd, error)
	GetByToken(token string) (*[]domain.Lgd, error)
}
