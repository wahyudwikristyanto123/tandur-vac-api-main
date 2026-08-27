package adapter

import (
	"tandur.com/src/adapter/out/mysql/inbasket/entity"
	"tandur.com/src/adapter/out/mysql/inbasket/repository"
	"tandur.com/src/adapter/out/mysql/inbasket/util"
	"tandur.com/src/domain"
)

type InbasketAdapter struct {
	repo *repository.InbasketBaseRepository
}

func NewInbasketAdapter(repo *repository.InbasketBaseRepository) InbasketBaseAdapter {
	return &InbasketAdapter{repo: repo}
}

func (adapter *InbasketAdapter) GetByToken(token string) (*[]domain.Inbasket, error) {
	inbasket, err := (*adapter.repo).GetFileByToken(token)
	if err != nil {
		return nil, err
	}
	events, err := (*adapter.repo).GetEventsByToken(token)
	if err != nil {
		return nil, err
	}
	emails, err := (*adapter.repo).GetEmailsByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.Inbasket = []domain.Inbasket{}
	data = append(data, util.EntityAdapterToDomain(*inbasket, *events, *emails))
	return &data, nil
}

func (adapter *InbasketAdapter) GetMailboxByToken(token string, statuses []string) (*[]domain.Mailbox, error) {
	results, err := (*adapter.repo).GetMailboxByToken(token, statuses)
	if err != nil {
		return nil, err
	}
	var data []domain.Mailbox = []domain.Mailbox{}
	i := 0
	for i < (len(*results)) {
		data = append(data, util.MailboxEntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *InbasketAdapter) SearchMailboxByToken(q string, token string, statuses []string) (*[]domain.Mailbox, error) {
	results, err := (*adapter.repo).SearchMailboxByToken(q, token, statuses)
	if err != nil {
		return nil, err
	}
	var data []domain.Mailbox = []domain.Mailbox{}
	i := 0
	for i < (len(*results)) {
		data = append(data, util.MailboxEntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *InbasketAdapter) GetMailboxById(id int64) (*domain.Mailbox, error) {
	result, err := (*adapter.repo).GetMailboxById(id)
	if err != nil {
		return nil, err
	}
	data := util.MailboxEntityAdapterToDomain((*result))
	return &data, nil
}

func (adapter *InbasketAdapter) GetMailboxByParentId(id int64) (*[]domain.Mailbox, error) {
	results, err := (*adapter.repo).GetMailboxByParentId(id)
	if err != nil {
		return nil, err
	}
	var data []domain.Mailbox = []domain.Mailbox{}
	i := 0
	for i < (len(*results)) {
		data = append(data, util.MailboxEntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *InbasketAdapter) GetMailboxByParentIdStatus(parentId int64, status string) (*domain.Mailbox, error) {
	result, err := (*adapter.repo).GetMailboxByParentIdStatus(parentId, status)
	if err != nil {
		return nil, err
	}
	data := util.MailboxEntityAdapterToDomain((*result))
	return &data, nil
}

func (adapter *InbasketAdapter) AddMailbox(data domain.Mailbox) (int64, error) {
	return (*adapter.repo).InsertMailbox(util.MailboxDomainToEntityAdapter(data))
}

func (adapter *InbasketAdapter) AddMultiMailbox(data []domain.Mailbox) error {
	var entries []entity.InbasketMailboxMySql
	for _, item := range data {
		entries = append(entries, util.MailboxDomainToEntityAdapter(item))
	}
	return (*adapter.repo).InsertMultiMailbox(entries)
}

func (adapter *InbasketAdapter) EditMailbox(data domain.Mailbox) error {
	return (*adapter.repo).UpdateMailbox(util.MailboxDomainToEntityAdapter(data))
}

func (adapter *InbasketAdapter) RemoveMailbox(id int64) error {
	return (*adapter.repo).DeleteMailbox(id)
}

func (adapter *InbasketAdapter) RemoveByParentIdStatus(id int64, status string) error {
	return (*adapter.repo).DeleteByParentIdStatus(id, status)
}
