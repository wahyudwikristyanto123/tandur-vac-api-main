package domain

import "time"

type Subtes struct {
	ID                 int64     `json:"id"`
	AssessmentCenterId int64     `json:"assessment_center_id"`
	Title              string    `json:"title"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	Duration           int64     `json:"duration"`
	Type               string    `json:"type"`
	ToolType           string    `json:"tool_type"`
	Index              int64     `json:"index"`
	URLProctoring      string    `json:"url_proctoring"`
	Status             string    `json:"status"`
}
