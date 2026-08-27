package repository

import (
	"tandur.com/src/adapter/out/mysql/inbasket/entity"
)

type InbasketBaseRepository interface {
	GetFileByToken(token string) (*entity.InbasketFileMySql, error)
	GetEventsByToken(token string) (*[]entity.InbasketEventMySql, error)
	GetEmailsByToken(token string) (*[]entity.InbasketEmailMySql, error)

	GetMailboxByToken(token string, statuses []string) (*[]entity.InbasketMailboxMySql, error)
	SearchMailboxByToken(q string, token string, statuses []string) (*[]entity.InbasketMailboxMySql, error)
	GetMailboxById(id int64) (*entity.InbasketMailboxMySql, error)
	GetMailboxByParentId(id int64) (*[]entity.InbasketMailboxMySql, error)
	GetMailboxByParentIdStatus(parentId int64, status string) (*entity.InbasketMailboxMySql, error)
	InsertMailbox(data entity.InbasketMailboxMySql) (int64, error)
	InsertMultiMailbox(data []entity.InbasketMailboxMySql) error
	UpdateMailbox(data entity.InbasketMailboxMySql) error
	DeleteMailbox(id int64) error
	DeleteByParentIdStatus(id int64, status string) error
}
