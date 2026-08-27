package domain

import (
	"time"
)

type Feedback struct {
	Id                 int64     `json:"id"`
	Token              string    `json:"token"`
	AssessmentCenterId int64     `json:"assessment_center_id"`
	Feedback           string    `json:"feedback"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
