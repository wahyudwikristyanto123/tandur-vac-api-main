package adapter

import (
	"tandur.com/src/adapter/out/mysql/subtes/repository"
	"tandur.com/src/adapter/out/mysql/subtes/util"
	"tandur.com/src/domain"
)

type SubtesAdapter struct {
	repo *repository.SubtesBaseRepository
}

func NewSubtesAdapter(repo *repository.SubtesBaseRepository) SubtesBaseAdapter {
	return &SubtesAdapter{repo: repo}
}

func (adapter *SubtesAdapter) GetAll(filter domain.Subtes) (*[]domain.Subtes, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.Subtes
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *SubtesAdapter) GetById(id int64) (*domain.Subtes, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *SubtesAdapter) GetResultById(id int64) (string, error) {
	return (*adapter.repo).GetResultById(id)
}

func (adapter *SubtesAdapter) GetByToken(token string) (*[]domain.Subtes, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.Subtes
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *SubtesAdapter) UpdateStatusById(id int64, status string) error {
	return (*adapter.repo).UpdateStatusById(id, status)
}

func (adapter *SubtesAdapter) UpdateResultById(id int64, result string) error {
	return (*adapter.repo).UpdateResultById(id, result)
}
