package adapter

import "tandur.com/src/domain"

type AssessorBaseAdapter interface {
	GetAll(filter domain.Assessor) (*[]domain.Assessor, error)
	GetById(id int64) (*domain.Assessor, error)
}
