package service

import (
	"fmt"
	"strings"
	"time"

	"tandur.com/src/adapter/out/mysql/inbasket/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
	"tandur.com/src/util/logger"
)

type InbasketService struct {
	adapter *adapter.InbasketBaseAdapter
}

func NewInbasketService(adapter *adapter.InbasketBaseAdapter) usecase.InbasketUseCase {
	return &InbasketService{
		adapter: adapter,
	}
}

func isArrayContain(haystack []string, needle string) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

func (service *InbasketService) GetByToken(q string, token string, statuses []string) (*[]domain.Inbasket, error) {
	inbasket, err := (*service.adapter).GetByToken(token)
	if err != nil {
		return nil, err
	}
	var mailbox *[]domain.Mailbox = &[]domain.Mailbox{}
	if isArrayContain(statuses, "UNREAD") {
		mailbox, err = (*service.adapter).SearchMailboxByToken(q, token, statuses)
		if err != nil {
			logger.WARN(err.Error())
			return nil, err
		}
		if len(*mailbox) < 1 {
			var messages []domain.Mailbox
			for i := 0; i < len(*inbasket); i++ {
				for j := 0; j < len((*inbasket)[i].Emails); j++ {
					email := (*inbasket)[i].Emails[j]
					messages = append(messages, domain.Mailbox{
						Token:       (*inbasket)[i].Token,
						Status:      "UNREAD",
						From:        email.From,
						Cc:          email.Cc,
						Subject:     email.Subject,
						Body:        email.Body,
						Attachments: email.Attachments,
						Date:        email.Date,
					})
				}
			}
			(*service.adapter).AddMultiMailbox(messages)
		}
	}
	mailbox, err = (*service.adapter).SearchMailboxByToken(q, token, statuses)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*inbasket); i++ {
		(*inbasket)[i].Mailbox = *mailbox
	}
	return inbasket, nil
}

func (service *InbasketService) GetById(id int64) (*domain.Mailbox, error) {
	data, err := (*service.adapter).GetMailboxById(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *InbasketService) GetByParentIdAndDraft(parentId int64) (*domain.Mailbox, error) {
	data, err := (*service.adapter).GetMailboxByParentIdStatus(parentId, "DRAFT")
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *InbasketService) Reply(parentId int64, message string) error {
	data, err := (*service.adapter).GetMailboxById(parentId)
	if err != nil {
		return err
	}
	(*service.adapter).EditMailbox(domain.Mailbox{
		Id:          parentId,
		Token:       data.Token,
		Status:      "REPLIED",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        data.Body,
		Attachments: data.Attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
	err = (*service.adapter).RemoveByParentIdStatus(parentId, "DRAFT")
	if err != nil {
		return err
	}
	_, err = (*service.adapter).AddMailbox(domain.Mailbox{
		Token:       data.Token,
		Status:      "SENT",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     fmt.Sprintf("Re: %s", data.Subject),
		Body:        message,
		Attachments: data.Attachments,
		Date:        time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ParentId:    parentId,
	})
	return err
}

func (service *InbasketService) FullReply(parentId int64, ccs []string, attachments []string, message string) error {
	data, err := (*service.adapter).GetMailboxById(parentId)
	if err != nil {
		return err
	}
	(*service.adapter).EditMailbox(domain.Mailbox{
		Id:          parentId,
		Token:       data.Token,
		Status:      "REPLIED",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        data.Body,
		Attachments: data.Attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
	err = (*service.adapter).RemoveByParentIdStatus(parentId, "DRAFT")
	if err != nil {
		return err
	}
	_, err = (*service.adapter).AddMailbox(domain.Mailbox{
		Token:       data.Token,
		Status:      "SENT",
		From:        data.From,
		Cc:          strings.Join(ccs, ","),
		Subject:     fmt.Sprintf("Re: %s", data.Subject),
		Body:        message,
		Attachments: attachments,
		Date:        time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ParentId:    parentId,
	})
	return err
}

func (service *InbasketService) ReplyAndDraft(parentId int64, message string) (int64, error) {
	data, err := (*service.adapter).GetMailboxById(parentId)
	if err != nil {
		return 0, err
	}
	(*service.adapter).EditMailbox(domain.Mailbox{
		Id:          parentId,
		Token:       data.Token,
		Status:      "READ",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        data.Body,
		Attachments: data.Attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
	drafts, err := (*service.adapter).GetMailboxByParentId(parentId)
	if err != nil {
		return 0, err
	}
	if len(*drafts) > 0 {
		(*service.adapter).EditMailbox(domain.Mailbox{
			Id:          (*drafts)[0].Id,
			Token:       data.Token,
			Status:      "DRAFT",
			From:        data.From,
			Cc:          data.Cc,
			Subject:     fmt.Sprintf("Re: %s", data.Subject),
			Body:        message,
			Attachments: data.Attachments,
			Date:        time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ParentId:    parentId,
		})
		return (*drafts)[0].Id, nil
	} else {
		return (*service.adapter).AddMailbox(domain.Mailbox{
			Token:       data.Token,
			Status:      "DRAFT",
			From:        data.From,
			Cc:          data.Cc,
			Subject:     fmt.Sprintf("Re: %s", data.Subject),
			Body:        message,
			Attachments: data.Attachments,
			Date:        time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ParentId:    parentId,
		})
	}
}

func (service *InbasketService) FullReplyAndDraft(parentId int64, ccs []string, attachments []string, message string) (int64, error) {
	data, err := (*service.adapter).GetMailboxById(parentId)
	if err != nil {
		return 0, err
	}
	(*service.adapter).EditMailbox(domain.Mailbox{
		Id:          parentId,
		Token:       data.Token,
		Status:      "READ",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        data.Body,
		Attachments: data.Attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
	drafts, err := (*service.adapter).GetMailboxByParentId(parentId)
	if err != nil {
		return 0, err
	}
	if len(*drafts) > 0 {
		(*service.adapter).EditMailbox(domain.Mailbox{
			Id:          (*drafts)[0].Id,
			Token:       data.Token,
			Status:      "DRAFT",
			From:        data.From,
			Cc:          strings.Join(ccs, ","),
			Subject:     fmt.Sprintf("Re: %s", data.Subject),
			Body:        message,
			Attachments: attachments,
			Date:        time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ParentId:    parentId,
		})
		return (*drafts)[0].Id, nil
	} else {
		return (*service.adapter).AddMailbox(domain.Mailbox{
			Token:       data.Token,
			Status:      "DRAFT",
			From:        data.From,
			Cc:          strings.Join(ccs, ","),
			Subject:     fmt.Sprintf("Re: %s", data.Subject),
			Body:        message,
			Attachments: attachments,
			Date:        time.Now(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			ParentId:    parentId,
		})
	}
}

func (service *InbasketService) UpdateAndDraft(id int64, message string) error {
	data, err := (*service.adapter).GetMailboxById(id)
	if err != nil {
		return err
	}
	return (*service.adapter).EditMailbox(domain.Mailbox{
		Id:          id,
		Token:       data.Token,
		Status:      "DRAFT",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        message,
		Attachments: data.Attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
}

func (service *InbasketService) FullUpdateAndDraft(id int64, ccs []string, attachments []string, message string) error {
	data, err := (*service.adapter).GetMailboxById(id)
	if err != nil {
		return err
	}
	return (*service.adapter).EditMailbox(domain.Mailbox{
		Id:          id,
		Token:       data.Token,
		Status:      "DRAFT",
		From:        data.From,
		Cc:          strings.Join(ccs, ","),
		Subject:     data.Subject,
		Body:        message,
		Attachments: attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
}

func (service *InbasketService) Create(data domain.Mailbox) error {
	_, err := (*service.adapter).AddMailbox(domain.Mailbox{
		Token:       data.Token,
		Status:      "SENT",
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        data.Body,
		Attachments: data.Attachments,
		Date:        time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		ParentId:    0,
	})
	return err
}

func (service *InbasketService) UpdateStatus(id int64, status string) error {
	data, err := (*service.adapter).GetMailboxById(id)
	if err != nil {
		return err
	}
	return (*service.adapter).EditMailbox(domain.Mailbox{
		Id:          id,
		Token:       data.Token,
		Status:      status,
		From:        data.From,
		Cc:          data.Cc,
		Subject:     data.Subject,
		Body:        data.Body,
		Attachments: data.Attachments,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   time.Now(),
		ParentId:    data.ParentId,
	})
}

func (service *InbasketService) DeleteMailbox(mailboxId int64) error {
	return (*service.adapter).RemoveMailbox(mailboxId)
}
