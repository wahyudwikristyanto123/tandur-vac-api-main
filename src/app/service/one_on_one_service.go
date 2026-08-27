package service

import (
	assessorAdapter "tandur.com/src/adapter/out/mysql/assessor/adapter"
	"tandur.com/src/adapter/out/mysql/one_on_one/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type OneOnOneService struct {
	adapter         *adapter.OneOnOneBaseAdapter
	assessorAdapter *assessorAdapter.AssessorBaseAdapter
}

func NewOneOnOneService(adapter *adapter.OneOnOneBaseAdapter, assessorAdapter *assessorAdapter.AssessorBaseAdapter) usecase.OneOnOneUseCase {
	return &OneOnOneService{
		adapter:         adapter,
		assessorAdapter: assessorAdapter,
	}
}

func (service *OneOnOneService) GetByToken(token string) (*[]domain.OneOnOne, error) {
	data, err := (*service.adapter).GetByToken(token)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*data); i++ {
		var finalAssessors []domain.Assessor = []domain.Assessor{}
		for j := 0; j < len((*data)[i].Assessors); j++ {
			finalAssessor, err := (*service.assessorAdapter).GetById((*data)[i].Assessors[j].Id)
			if err != nil {
				continue
			}
			finalAssessors = append(finalAssessors, *finalAssessor)
		}
		(*data)[i].Assessors = finalAssessors
	}
	return data, nil
}
