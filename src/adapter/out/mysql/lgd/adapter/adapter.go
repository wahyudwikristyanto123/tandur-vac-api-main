package adapter

import (
	"tandur.com/src/adapter/out/mysql/lgd/repository"
	"tandur.com/src/adapter/out/mysql/lgd/util"
	"tandur.com/src/domain"
)

type LgdAdapter struct {
	repo *repository.LgdBaseRepository
}

func NewLgdAdapter(repo *repository.LgdBaseRepository) LgdBaseAdapter {
	return &LgdAdapter{repo: repo}
}

func (adapter *LgdAdapter) GetAll(filter domain.Lgd) (*[]domain.Lgd, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.Lgd
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *LgdAdapter) GetById(id int64) (*domain.Lgd, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *LgdAdapter) GetByToken(token string) (*[]domain.Lgd, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.Lgd
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}
