package domain

import "time"

type RolePlay struct {
	Token         string     `json:"token"`
	MeetingUrl    string     `json:"meeting_url"`
	StartDate     time.Time  `json:"start_date"`
	EndDate       time.Time  `json:"end_date"`
	Type          string     `json:"type"`
	Assessors     []Assessor `json:"assessors"`
	Participants  []Assessor `json:"participants"`
	Roleplayers   []Assessor `json:"roleplayers"`
	RelatedAnswer string     `json:"related_answer"`
}
