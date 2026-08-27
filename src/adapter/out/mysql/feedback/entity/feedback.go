package entity

import (
	"database/sql"
)

type FeedbackMySql struct {
	Id                 int64          `json:"id"`
	Token              sql.NullString `json:"token"`
	AssessmentCenterId int64          `json:"trx_assessment_center_id"`
	Feedback           sql.NullString `json:"feedback"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
}
