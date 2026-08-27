package entity

import (
	"database/sql"
)

type CvMySql struct {
	Id                 int64          `json:"id"`
	Token              sql.NullString `json:"token"`
	AssessmentCenterId int64          `json:"trx_assessment_center_id"`
	Page1              sql.NullString `json:"page1"`
	Page2              sql.NullString `json:"page2"`
	Page3              sql.NullString `json:"page3"`
	Page4              sql.NullString `json:"page4"`
	Page5              sql.NullString `json:"page5"`
	Page6              sql.NullString `json:"page6"`
	Page7              sql.NullString `json:"page7"`
	Page8              sql.NullString `json:"page8"`
	Page9              sql.NullString `json:"page9"`
	Page10             sql.NullString `json:"page10"`
	Page11             sql.NullString `json:"page11"`
	Page12             sql.NullString `json:"page12"`
	Page13             sql.NullString `json:"page13"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
}

type CvViewMySql struct {
	Token              sql.NullString `json:"token"`
	CmdId              int64          `json:"cmd_id"`
	Id                 int64          `json:"id"`
	AssessmentCenterId int64          `json:"trx_assessment_center_id"`
	Page1              sql.NullString `json:"page1"`
	Page2              sql.NullString `json:"page2"`
	Page3              sql.NullString `json:"page3"`
	Page4              sql.NullString `json:"page4"`
	Page5              sql.NullString `json:"page5"`
	Page6              sql.NullString `json:"page6"`
	Page7              sql.NullString `json:"page7"`
	Page8              sql.NullString `json:"page8"`
	Page9              sql.NullString `json:"page9"`
	Page10             sql.NullString `json:"page10"`
	Page11             sql.NullString `json:"page11"`
	Page12             sql.NullString `json:"page12"`
	Page13             sql.NullString `json:"page13"`
	CreatedAt          sql.NullTime   `json:"created_at"`
	UpdatedAt          sql.NullTime   `json:"updated_at"`
}
