package service

import (
	"tandur.com/src/adapter/out/mysql/qna/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type QnaService struct {
	adapter *adapter.QnaBaseAdapter
}

func NewQnaService(adapter *adapter.QnaBaseAdapter) usecase.QnaUseCase {
	return &QnaService{
		adapter: adapter,
	}
}

func (service *QnaService) GetById(id int64) (*domain.Qna, error) {
	data, err := (*service.adapter).GetById(id)
	if err != nil {
		return nil, err
	}
	details, err := (*service.adapter).GetDetailsById((*data).Id)
	if err != nil {
		return nil, err
	}
	(*data).Details = *details
	return data, nil
}

func (service *QnaService) GetByToken(token string) (*[]domain.Qna, error) {
	data, err := (*service.adapter).GetByToken(token)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*data); i++ {
		details, err := (*service.adapter).GetDetailsById((*data)[i].Id)
		if err != nil {
			return nil, err
		}
		(*data)[i].Details = *details
	}
	return data, nil
}

func (service *QnaService) UpsertResult(data domain.QnaResultRequest) error {
	return (*service.adapter).UpsertResult(data)
}

func (service *QnaService) UpsertResults(data []domain.QnaResultRequest) error {
	for i := 0; i < len(data); i++ {
		(*service.adapter).UpsertResult(data[i])
	}
	return nil
}

func (service *QnaService) GetResultsByToken(token string) (*[]domain.QnaResultRequest, error) {
	return (*service.adapter).GetResultsByToken(token)
}
