package adapter

import (
	"tandur.com/src/adapter/out/mysql/qna/repository"
	"tandur.com/src/adapter/out/mysql/qna/util"
	"tandur.com/src/domain"
)

type QnaAdapter struct {
	repo *repository.QnaBaseRepository
}

func NewQnaAdapter(repo *repository.QnaBaseRepository) QnaBaseAdapter {
	return &QnaAdapter{repo: repo}
}

func (adapter *QnaAdapter) GetAll(filter domain.Qna) (*[]domain.Qna, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.Qna
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *QnaAdapter) GetDetailsById(id int64) (*[]domain.QnaDetail, error) {
	results, err := (*adapter.repo).GetDetailsById(id)
	if err != nil {
		return nil, err
	}
	var data []domain.QnaDetail
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterDetailToDomainDetail((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *QnaAdapter) GetById(id int64) (*domain.Qna, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *QnaAdapter) GetByToken(token string) (*[]domain.Qna, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.Qna
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *QnaAdapter) UpsertResult(data domain.QnaResultRequest) error {
	return (*adapter.repo).UpsertResult(util.DomainToEntityAdapterResult(data))
}

func (adapter *QnaAdapter) GetResultsByToken(token string) (*[]domain.QnaResultRequest, error) {
	results, err := (*adapter.repo).GetResultsByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.QnaResultRequest
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterResultToDomain((*results)[i]))
		i++
	}
	return &data, nil
}
