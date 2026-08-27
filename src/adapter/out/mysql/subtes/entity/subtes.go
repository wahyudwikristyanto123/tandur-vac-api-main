package entity

import (
	"database/sql"
)

type SubtesMySql struct {
	Token              sql.NullString `json:"token"`
	ID                 int64          `json:"id"`
	UserName           sql.NullString `json:"user_name"`
	AssessmentCenterId int64          `json:"assessment_center_id"`
	Title              string         `json:"title"`
	StartDate          sql.NullTime   `json:"start_date"`
	EndDate            sql.NullTime   `json:"end_date"`
	ToolsType          string         `json:"tools_type"`
	Type               string         `json:"type"`
	OrderSchedule      sql.NullInt64  `json:"order_schedule"`
	URLProctoring      sql.NullString `json:"url_proctoring"`
	Status             sql.NullString `json:"status"`
	Duration           int64          `json:"duration"`
}
