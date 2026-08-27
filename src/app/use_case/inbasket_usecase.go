package usecase

import "tandur.com/src/domain"

type InbasketUseCase interface {
	GetByToken(q string, token string, statuses []string) (*[]domain.Inbasket, error)
	GetById(id int64) (*domain.Mailbox, error)
	GetByParentIdAndDraft(parentId int64) (*domain.Mailbox, error)
	Reply(parentId int64, message string) error
	FullReply(parentId int64, ccs []string, attachments []string, message string) error
	ReplyAndDraft(parentId int64, message string) (int64, error)
	FullReplyAndDraft(parentId int64, ccs []string, attachments []string, message string) (int64, error)
	UpdateAndDraft(id int64, message string) error
	FullUpdateAndDraft(id int64, ccs []string, attachments []string, message string) error
	Create(data domain.Mailbox) error
	UpdateStatus(mailboxId int64, status string) error
	DeleteMailbox(mailboxId int64) error
}
