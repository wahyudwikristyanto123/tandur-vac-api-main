package repository

import (
	"tandur.com/src/adapter/out/mysql/problem_analysis/entity"
	"tandur.com/src/domain"
)

type ProblemAnalysisBaseRepository interface {
	GetAll(filter domain.ProblemAnalysis) (*[]entity.ProblemAnalysisMySql, error)
	GetById(id int64) (*entity.ProblemAnalysisMySql, error)
	GetByToken(token string) (*[]entity.ProblemAnalysisMySql, error)
}
