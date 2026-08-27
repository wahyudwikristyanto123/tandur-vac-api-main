package entity

import (
	"database/sql"
)

type ProblemAnalysisMySql struct {
	Token              sql.NullString `json:"token"`
	ToolsType          string         `json:"tools_type"`
	InstructionFile    sql.NullString `json:"instruction_file"`
	Instruction        sql.NullString `json:"instruction"`
	CompanyProfileFile sql.NullString `json:"company_profile_file"`
	CompanyProfile     sql.NullString `json:"company_profile"`
	Question           string         `json:"question"`
	StartDate          sql.NullTime   `json:"start_date"`
	EndDate            sql.NullTime   `json:"end_date"`
}
