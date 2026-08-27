package adapter

import (
	"tandur.com/src/adapter/out/mysql/assessor/repository"
	"tandur.com/src/adapter/out/mysql/assessor/util"
	"tandur.com/src/domain"
)

type AssessorAdapter struct {
	repo *repository.AssessorBaseRepository
}

func NewAssessorAdapter(repo *repository.AssessorBaseRepository) AssessorBaseAdapter {
	return &AssessorAdapter{repo: repo}
}

func (adapter *AssessorAdapter) GetAll(filter domain.Assessor) (*[]domain.Assessor, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.Assessor
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *AssessorAdapter) GetById(id int64) (*domain.Assessor, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}
