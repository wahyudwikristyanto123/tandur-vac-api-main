package service

import (
	assessorAdapter "tandur.com/src/adapter/out/mysql/assessor/adapter"
	"tandur.com/src/adapter/out/mysql/roleplay/adapter"
	subtesAdapter "tandur.com/src/adapter/out/mysql/subtes/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type RoleplayService struct {
	adapter         *adapter.RoleplayBaseAdapter
	assessorAdapter *assessorAdapter.AssessorBaseAdapter
	subtesAdapter   *subtesAdapter.SubtesBaseAdapter
}

func NewRoleplayService(adapter *adapter.RoleplayBaseAdapter, assessorAdapter *assessorAdapter.AssessorBaseAdapter, subtesAdapter *subtesAdapter.SubtesBaseAdapter) usecase.RoleplayUseCase {
	return &RoleplayService{
		adapter:         adapter,
		assessorAdapter: assessorAdapter,
		subtesAdapter:   subtesAdapter,
	}
}

func (service *RoleplayService) getRelatedAnswer(token string) (string, error) {
	var result = ""
	tests, err := (*service.subtesAdapter).GetByToken(token)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(*tests); i++ {
		if (*tests)[i].ToolType == "pa" {
			result, err = (*service.subtesAdapter).GetResultById((*tests)[i].ID)
			if err != nil {
				return "", err
			}
		}
	}
	return result, nil
}

func (service *RoleplayService) GetByToken(token string) (*[]domain.RolePlay, error) {
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

		var finalParticipants []domain.Assessor = []domain.Assessor{}
		for j := 0; j < len((*data)[i].Participants); j++ {
			finalParticipant, err := (*service.assessorAdapter).GetById((*data)[i].Participants[j].Id)
			if err != nil {
				continue
			}
			finalParticipants = append(finalParticipants, *finalParticipant)
		}
		(*data)[i].Participants = finalParticipants

		var finalRoleplayers []domain.Assessor = []domain.Assessor{}
		for j := 0; j < len((*data)[i].Roleplayers); j++ {
			finalRoleplayer, err := (*service.assessorAdapter).GetById((*data)[i].Roleplayers[j].Id)
			if err != nil {
				continue
			}
			finalRoleplayers = append(finalRoleplayers, *finalRoleplayer)
		}
		(*data)[i].Roleplayers = finalRoleplayers

		paResult, err := service.getRelatedAnswer(token)
		if err != nil {
			return nil, err
		}
		(*data)[i].RelatedAnswer = paResult
	}
	return data, nil
}
