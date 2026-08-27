package adapter

import (
	"tandur.com/src/domain"
)

type CvBaseAdapter interface {
	GetCvByToken(token string) (*[]domain.Cv, error)
	GetCvById(id int64) (*domain.Cv, error)
	AddCv(data domain.Cv) (int64, error)
	UpdateCv(data domain.Cv) error
	RemoveCv(id int64) error
}
