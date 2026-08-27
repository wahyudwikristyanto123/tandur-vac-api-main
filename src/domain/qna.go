package domain

import (
	"time"
)

type Qna struct {
	Id             int64       `json:"id"`
	IdQna          int64       `json:"id_qna"`
	Token          string      `json:"token"`
	InstructionUrl string      `json:"instruction_url"`
	Instruction    string      `json:"instruction"`
	Type           string      `json:"type"`
	Details        []QnaDetail `json:"details"`
	StartDate      time.Time   `json:"start_date"`
	EndDate        time.Time   `json:"end_date"`
}

type QnaDetail struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parent_id"`
	Title    string `json:"title"`
	Question string `json:"question"`
}

type QnaResultRequest struct {
	QuestionId int64  `json:"question_id"`
	QnaId      int64  `json:"qna_id"`
	Token      string `json:"token"`
	Result     string `json:"result"`
}
