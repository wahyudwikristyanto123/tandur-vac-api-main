package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/problem_analysis/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type ProblemAnalysisRepository struct {
	tableName string
}

func NewProblemAnalysisRepository() ProblemAnalysisBaseRepository {
	return &ProblemAnalysisRepository{
		tableName: "V_GET_PA",
	}
}

func (repo *ProblemAnalysisRepository) GetAll(filter domain.ProblemAnalysis) (*[]entity.ProblemAnalysisMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, company_profile_file, company_profile, question, start_date, end_date FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.ProblemAnalysisMySql
	for results.Next() {
		var res entity.ProblemAnalysisMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.CompanyProfileFile, &res.CompanyProfile, &res.Question, &res.StartDate, &res.EndDate)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *ProblemAnalysisRepository) GetById(id int64) (*entity.ProblemAnalysisMySql, error) {
	db := util.GetMySQL()
	var res entity.ProblemAnalysisMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, company_profile_file, company_profile, question, start_date, end_date FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.CompanyProfileFile, &res.CompanyProfile, &res.Question, &res.StartDate, &res.EndDate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *ProblemAnalysisRepository) GetByToken(token string) (*[]entity.ProblemAnalysisMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, company_profile_file, company_profile, question, start_date, end_date FROM %s WHERE token = ?", repo.tableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.ProblemAnalysisMySql
	for results.Next() {
		var res entity.ProblemAnalysisMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.CompanyProfileFile, &res.CompanyProfile, &res.Question, &res.StartDate, &res.EndDate)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}
