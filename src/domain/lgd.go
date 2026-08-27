package domain

import "time"

type Lgd struct {
	Token          string     `json:"token"`
	MeetingUrl     string     `json:"meeting_url"`
	Introduction   string     `json:"introduction"`
	Instruction    string     `json:"instruction"`
	InstructionUrl string     `json:"instruction_url"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        time.Time  `json:"end_date"`
	Type           string     `json:"type"`
	Assessors      []Assessor `json:"assessors"`
	Participants   []Assessor `json:"participants"`
	RelatedAnswer  string     `json:"related_answer"`
}
