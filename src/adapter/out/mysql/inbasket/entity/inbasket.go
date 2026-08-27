package entity

import (
	"database/sql"
	"time"
)

type InbasketFileMySql struct {
	Token                     sql.NullString `json:"token"`
	ToolsType                 string         `json:"tools_type"`
	InstructionFile           sql.NullString `json:"instruction_file"`
	Instruction               sql.NullString `json:"instruction"`
	CompanyProfileFile        sql.NullString `json:"company_profile_file"`
	OrganizationStructureFile sql.NullString `json:"organization_structure_file"`
	StartDate                 sql.NullTime   `json:"start_date"`
	EndDate                   sql.NullTime   `json:"end_date"`
}

type InbasketEventMySql struct {
	Token            sql.NullString `json:"token"`
	ToolsType        string         `json:"tools_type"`
	EventId          int64          `json:"event_id"`
	EventTitle       string         `json:"event_title"`
	EventDate        time.Time      `json:"event_date"`
	EventDescription string         `json:"event_description"`
	StartDate        sql.NullTime   `json:"start_date"`
	EndDate          sql.NullTime   `json:"end_date"`
}

type InbasketEmailMySql struct {
	Token         sql.NullString `json:"token"`
	ToolsType     string         `json:"tools_type"`
	EmailId       int64          `json:"email_id"`
	EmailFrom     string         `json:"email_from"`
	EmailCc       string         `json:"email_cc"`
	EmailSubject  string         `json:"email_subject"`
	EmailBody     string         `json:"email_body"`
	EmailSendDate sql.NullTime   `json:"email_send_date"`
	Attachment1   sql.NullString `json:"attachment_1"`
	Attachment2   sql.NullString `json:"attachment_2"`
	Attachment3   sql.NullString `json:"attachment_3"`
	Attachment4   sql.NullString `json:"attachment_4"`
	StartDate     sql.NullTime   `json:"start_date"`
	EndDate       sql.NullTime   `json:"end_date"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
}

type InbasketMailboxMySql struct {
	Id            int64          `json:"email_id"`
	Token         sql.NullString `json:"token"`
	EmailFrom     string         `json:"email_from"`
	EmailCc       string         `json:"email_cc"`
	EmailSubject  string         `json:"email_subject"`
	EmailBody     string         `json:"email_body"`
	EmailSendDate sql.NullTime   `json:"email_send_date"`
	Attachment1   sql.NullString `json:"attachment_1"`
	Attachment2   sql.NullString `json:"attachment_2"`
	Attachment3   sql.NullString `json:"attachment_3"`
	Attachment4   sql.NullString `json:"attachment_4"`
	Status        string         `json:"status"`
	ParentId      sql.NullInt64  `json:"parent_id"`
	CreatedAt     sql.NullTime   `json:"created_at"`
	UpdatedAt     sql.NullTime   `json:"updated_at"`
}
