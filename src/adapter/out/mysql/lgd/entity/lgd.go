package entity

import (
	"database/sql"
)

type LgdMySql struct {
	Token                  sql.NullString `json:"token"`
	ToolsType              string         `json:"tools_type"`
	InstructionFile        sql.NullString `json:"instruction_file"`
	Instruction            string         `json:"instruction"`
	MeetingUrl             sql.NullString `json:"meeting_url"`
	ParticipantInstruction string         `json:"participant_instruction"`
	StartDate              sql.NullTime   `json:"start_date"`
	EndDate                sql.NullTime   `json:"end_date"`
	Assessor1Id            int64          `json:"assessor_1_id"`
	Assessor2Id            sql.NullInt64  `json:"assessor_2_id"`
	Assessor3Id            sql.NullInt64  `json:"assessor_3_id"`
	Assessor4Id            sql.NullInt64  `json:"assessor_4_id"`
	Asesi                  sql.NullString `json:"asesi"`
}
