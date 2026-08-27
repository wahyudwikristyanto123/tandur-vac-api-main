package adapter

import "tandur.com/src/domain"

type InbasketBaseAdapter interface {
	GetByToken(token string) (*[]domain.Inbasket, error)
	GetMailboxByToken(token string, statuses []string) (*[]domain.Mailbox, error)
	SearchMailboxByToken(q string, token string, statuses []string) (*[]domain.Mailbox, error)
	GetMailboxById(id int64) (*domain.Mailbox, error)
	GetMailboxByParentId(id int64) (*[]domain.Mailbox, error)
	GetMailboxByParentIdStatus(parentId int64, status string) (*domain.Mailbox, error)
	AddMailbox(data domain.Mailbox) (int64, error)
	AddMultiMailbox(data []domain.Mailbox) error
	EditMailbox(data domain.Mailbox) error
	RemoveMailbox(id int64) error
	RemoveByParentIdStatus(id int64, status string) error
}
