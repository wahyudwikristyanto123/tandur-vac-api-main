package service

import (
	"tandur.com/src/adapter/out/mysql/subtes/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type SubtesService struct {
	adapter *adapter.SubtesBaseAdapter
}

func NewSubtesService(adapter *adapter.SubtesBaseAdapter) usecase.SubtesUseCase {
	return &SubtesService{
		adapter: adapter,
	}
}

func (service *SubtesService) GetByToken(token string) (*[]domain.Subtes, error) {
	data, err := (*service.adapter).GetByToken(token)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *SubtesService) GetResultById(id int64) (string, error) {
	data, err := (*service.adapter).GetResultById(id)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (service *SubtesService) UpdateStatusById(id int64, status string) error {
	return (*service.adapter).UpdateStatusById(id, status)
}

func (service *SubtesService) SubmitResult(id int64, result string) error {
	return (*service.adapter).UpdateResultById(id, result)
}
