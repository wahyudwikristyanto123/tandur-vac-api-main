package domain

import "time"

type ProblemAnalysis struct {
	Token             string    `json:"token"`
	Question          string    `json:"question"`
	Instruction       string    `json:"instruction"`
	InstructionUrl    string    `json:"instruction_url"`
	CompanyProfileUrl string    `json:"company_profile_url"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date"`
	Type              string    `json:"type"`
}
