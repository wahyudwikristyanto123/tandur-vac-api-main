package adapter

import (
	"tandur.com/src/adapter/out/mysql/roleplay/repository"
	"tandur.com/src/adapter/out/mysql/roleplay/util"
	"tandur.com/src/domain"
)

type RoleplayAdapter struct {
	repo *repository.RoleplayBaseRepository
}

func NewRoleplayAdapter(repo *repository.RoleplayBaseRepository) RoleplayBaseAdapter {
	return &RoleplayAdapter{repo: repo}
}

func (adapter *RoleplayAdapter) GetAll(filter domain.RolePlay) (*[]domain.RolePlay, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.RolePlay
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *RoleplayAdapter) GetById(id int64) (*domain.RolePlay, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *RoleplayAdapter) GetByToken(token string) (*[]domain.RolePlay, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.RolePlay
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}
