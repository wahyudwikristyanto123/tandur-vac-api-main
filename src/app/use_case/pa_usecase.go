package usecase

import "tandur.com/src/domain"

type ProblemAnalysisUseCase interface {
	GetByToken(token string) (*[]domain.ProblemAnalysis, error)
}
