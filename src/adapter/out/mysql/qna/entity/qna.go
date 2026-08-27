package entity

import (
	"database/sql"
)

type QnaMySql struct {
	Id              int64          `json:"qna_id"`
	IdQna           int64          `json:"detail_id"`
	Token           sql.NullString `json:"token"`
	InstructionFile sql.NullString `json:"instruction_file"`
	Instruction     string         `json:"instruction"`
	TypeQna         sql.NullString `json:"type_qna"`
	StartDate       sql.NullTime   `json:"start_date"`
	EndDate         sql.NullTime   `json:"end_date"`
}

type QnaDetailMySql struct {
	Id                    int64        `json:"token"`
	TrxAcQuestionAnswerId int64        `json:"trx_ac_question_answer_id"`
	HeadQuestion          string       `json:"head_question"`
	Question              string       `json:"question"`
	CreatedAt             sql.NullTime `json:"created_at"`
	UpdatedAt             sql.NullTime `json:"updated_at"`
}

type QnaResultRequestMysql struct {
	QuestionId int64          `json:"question_id"`
	QnaId      int64          `json:"qna_id"`
	Token      string         `json:"token"`
	Result     sql.NullString `json:"result"`
}
