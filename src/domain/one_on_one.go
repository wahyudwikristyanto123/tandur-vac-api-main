package domain

import "time"

type OneOnOne struct {
	Token      string     `json:"token"`
	MeetingUrl string     `json:"meeting_url"`
	StartDate  time.Time  `json:"start_date"`
	EndDate    time.Time  `json:"end_date"`
	Type       string     `json:"type"`
	Assessors  []Assessor `json:"assessors"`
}
