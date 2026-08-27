package adapter

import (
	"tandur.com/src/adapter/out/mysql/cv/repository"
	"tandur.com/src/adapter/out/mysql/cv/util"
	"tandur.com/src/domain"
)

type CvAdapter struct {
	repo *repository.CvBaseRepository
}

func NewCvAdapter(repo *repository.CvBaseRepository) CvBaseAdapter {
	return &CvAdapter{repo: repo}
}

func (adapter *CvAdapter) GetCvByToken(token string) (*[]domain.Cv, error) {
	results, err := (*adapter.repo).GetCvByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.Cv = []domain.Cv{}
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *CvAdapter) GetCvById(id int64) (*domain.Cv, error) {
	result, err := (*adapter.repo).GetCvById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain((*result))
	return &data, nil
}

func (adapter *CvAdapter) AddCv(data domain.Cv) (int64, error) {
	return (*adapter.repo).InsertCv(util.DomainToEntityAdapter(data))
}

func (adapter *CvAdapter) UpdateCv(data domain.Cv) error {
	return (*adapter.repo).UpdateCv(util.DomainToEntityAdapter(data))
}

func (adapter *CvAdapter) RemoveCv(id int64) error {
	return (*adapter.repo).DeleteCv(id)
}
