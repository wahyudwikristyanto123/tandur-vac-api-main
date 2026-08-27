package service

import (
	"tandur.com/src/adapter/out/mysql/user/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type UserService struct {
	adapter *adapter.UserBaseAdapter
}

func NewUserService(adapter *adapter.UserBaseAdapter) usecase.UserUseCase {
	return &UserService{
		adapter: adapter,
	}
}

func (service *UserService) GetUserByToken(token string) (*[]domain.User, error) {
	data, err := (*service.adapter).GetByToken(token)
	if err != nil {
		return nil, err
	}
	return data, nil
}
