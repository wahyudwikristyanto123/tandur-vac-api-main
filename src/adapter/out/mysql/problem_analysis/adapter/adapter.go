package adapter

import (
	"tandur.com/src/adapter/out/mysql/problem_analysis/repository"
	"tandur.com/src/adapter/out/mysql/problem_analysis/util"
	"tandur.com/src/domain"
)

type ProblemAnalysisAdapter struct {
	repo *repository.ProblemAnalysisBaseRepository
}

func NewProblemAnalysisAdapter(repo *repository.ProblemAnalysisBaseRepository) ProblemAnalysisBaseAdapter {
	return &ProblemAnalysisAdapter{repo: repo}
}

func (adapter *ProblemAnalysisAdapter) GetAll(filter domain.ProblemAnalysis) (*[]domain.ProblemAnalysis, error) {
	results, err := (*adapter.repo).GetAll(filter)
	if err != nil {
		return nil, err
	}
	var data []domain.ProblemAnalysis
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *ProblemAnalysisAdapter) GetById(id int64) (*domain.ProblemAnalysis, error) {
	res, err := (*adapter.repo).GetById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain(*res)
	return &data, nil
}

func (adapter *ProblemAnalysisAdapter) GetByToken(token string) (*[]domain.ProblemAnalysis, error) {
	results, err := (*adapter.repo).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.ProblemAnalysis
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}
