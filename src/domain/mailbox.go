package domain

import "time"

type Mailbox struct {
	Id          int64     `json:"id"`
	Token       string    `json:"token"`
	Status      string    `json:"status"`
	From        string    `json:"from"`
	Cc          string    `json:"cc"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	Attachments []string  `json:"attachments"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ParentId    int64     `json:"parent_id"`
}
