package service

import (
	"tandur.com/src/adapter/out/mysql/problem_analysis/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type ProblemAnalysisService struct {
	adapter *adapter.ProblemAnalysisBaseAdapter
}

func NewProblemAnalysisService(adapter *adapter.ProblemAnalysisBaseAdapter) usecase.ProblemAnalysisUseCase {
	return &ProblemAnalysisService{
		adapter: adapter,
	}
}

func (service *ProblemAnalysisService) GetByToken(token string) (*[]domain.ProblemAnalysis, error) {
	data, err := (*service.adapter).GetByToken(token)
	if err != nil {
		return nil, err
	}
	return data, nil
	// var final []domain.ProblemAnalysis
	// for i := 0; i < len(*data); i++ {
	// 	pa := (*data)[0]
	// 	splitted := strings.Split(pa.CompanyProfileUrl, ".")
	// 	var joined []string
	// 	for j := 0; j < len(splitted)-1; j++ {
	// 		joined = append(joined, splitted[j])
	// 	}
	// 	finalUrl := strings.Join(joined, ".")
	// 	pa.CompanyProfileUrl = finalUrl
	// 	final = append(final, pa)
	// }
	// return &final, nil
}
