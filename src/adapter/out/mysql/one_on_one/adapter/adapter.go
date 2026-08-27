package adapter

import (
	"tandur.com/src/adapter/out/mysql/one_on_one/repository"
	"tandur.com/src/adapter/out/mysql/one_on_one/util"
	"tandur.com/src/domain"
)

type OneOnOneAdapter struct {
	repo *repository.OneOnOneBaseRepository
}

func NewOneOnOneAdapter(repo *repository.OneOnOneBaseRepository) OneOnOneBaseAdapter {
	return &OneOnOneAdapter{repo: repo}
}

func (adapter *OneOnOneAdapter) GetAll(filter domain.OneOnOne) (*[]domain.OneOnOne, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.OneOnOne
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *OneOnOneAdapter) GetById(id int64) (*domain.OneOnOne, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *OneOnOneAdapter) GetByToken(token string) (*[]domain.OneOnOne, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.OneOnOne
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}
