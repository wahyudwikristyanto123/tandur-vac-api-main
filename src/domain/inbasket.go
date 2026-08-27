package domain

import "time"

type Inbasket struct {
	Token             string          `json:"token"`
	InstructionUrl    string          `json:"instruction_url"`
	Instruction       string          `json:"instruction"`
	CompanyProfileUrl string          `json:"company_profile_url"`
	OrganizationUrl   string          `json:"organization_url"`
	Events            []InbasketEvent `json:"events"`
	Emails            []InbasketEmail `json:"emails"`
	Mailbox           []Mailbox       `json:"mailbox"`
	StartDate         time.Time       `json:"start_date"`
	EndDate           time.Time       `json:"end_date"`
	Type              string          `json:"type"`
}

type InbasketEvent struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

type InbasketEmail struct {
	Id          int64     `json:"id"`
	From        string    `json:"from"`
	Cc          string    `json:"cc"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	Attachments []string  `json:"attachments"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
