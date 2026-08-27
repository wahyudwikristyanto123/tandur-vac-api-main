package adapter

import "tandur.com/src/domain"

type ProblemAnalysisBaseAdapter interface {
	GetAll(filter domain.ProblemAnalysis) (*[]domain.ProblemAnalysis, error)
	GetById(id int64) (*domain.ProblemAnalysis, error)
	GetByToken(token string) (*[]domain.ProblemAnalysis, error)
}
