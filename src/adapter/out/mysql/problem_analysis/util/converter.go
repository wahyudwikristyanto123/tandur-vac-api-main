package util

import (
	"tandur.com/src/adapter/out/mysql/problem_analysis/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(data entity.ProblemAnalysisMySql) domain.ProblemAnalysis {
	return domain.ProblemAnalysis{
		Token:             data.Token.String,
		Question:          data.Question,
		Instruction:       data.Instruction.String,
		InstructionUrl:    data.InstructionFile.String,
		CompanyProfileUrl: data.CompanyProfileFile.String,
		StartDate:         data.StartDate.Time,
		EndDate:           data.EndDate.Time,
		Type:              data.ToolsType,
	}
}
