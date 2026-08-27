package adapter

import (
	"tandur.com/src/adapter/out/mysql/user/repository"
	"tandur.com/src/adapter/out/mysql/user/util"
	"tandur.com/src/domain"
)

type UserAdapter struct {
	repo *repository.UserBaseRepository
}

func NewUserAdapter(repo *repository.UserBaseRepository) UserBaseAdapter {
	return &UserAdapter{repo: repo}
}

func (adapter *UserAdapter) GetAll(filter domain.User) (*[]domain.User, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.User
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *UserAdapter) GetById(id int64) (*domain.User, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *UserAdapter) GetByToken(token string) (*[]domain.User, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.User
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}
