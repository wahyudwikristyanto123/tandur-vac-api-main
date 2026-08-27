package util

import (
	"database/sql"

	"tandur.com/src/adapter/out/mysql/inbasket/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(file entity.InbasketFileMySql, events []entity.InbasketEventMySql, emails []entity.InbasketEmailMySql) domain.Inbasket {
	var finalEvents []domain.InbasketEvent = []domain.InbasketEvent{}
	for i := 0; i < len(events); i++ {
		finalEvents = append(finalEvents, domain.InbasketEvent{
			Id:          events[i].EventId,
			Title:       events[i].EventTitle,
			Date:        events[i].EventDate,
			Description: events[i].EventDescription,
		})
	}
	var finalEmails []domain.InbasketEmail = []domain.InbasketEmail{}
	for i := 0; i < len(emails); i++ {
		var attachments []string = []string{}
		if emails[i].Attachment1.Valid && emails[i].Attachment1.String != "" {
			attachments = append(attachments, emails[i].Attachment1.String)
		}
		if emails[i].Attachment2.Valid && emails[i].Attachment2.String != "" {
			attachments = append(attachments, emails[i].Attachment2.String)
		}
		if emails[i].Attachment3.Valid && emails[i].Attachment3.String != "" {
			attachments = append(attachments, emails[i].Attachment3.String)
		}
		if emails[i].Attachment4.Valid && emails[i].Attachment4.String != "" {
			attachments = append(attachments, emails[i].Attachment4.String)
		}

		finalEmails = append(finalEmails, domain.InbasketEmail{
			Id:          emails[i].EmailId,
			From:        emails[i].EmailFrom,
			Cc:          emails[i].EmailCc,
			Subject:     emails[i].EmailSubject,
			Body:        emails[i].EmailBody,
			Attachments: attachments,
			Date:        emails[i].EmailSendDate.Time,
			CreatedAt:   emails[i].CreatedAt.Time,
			UpdatedAt:   emails[i].UpdatedAt.Time,
		})
	}
	return domain.Inbasket{
		Token:             file.Token.String,
		Instruction:       file.Instruction.String,
		InstructionUrl:    file.InstructionFile.String,
		CompanyProfileUrl: file.CompanyProfileFile.String,
		OrganizationUrl:   file.OrganizationStructureFile.String,
		Events:            finalEvents,
		Emails:            finalEmails,
		StartDate:         file.StartDate.Time,
		EndDate:           file.EndDate.Time,
		Type:              file.ToolsType,
	}
}

func isValid(val string) bool {
	return val != ""
}

func MailboxEntityAdapterToDomain(data entity.InbasketMailboxMySql) domain.Mailbox {
	var attachments []string = []string{}
	if data.Attachment1.Valid && data.Attachment1.String != "" {
		attachments = append(attachments, data.Attachment1.String)
	}
	if data.Attachment2.Valid && data.Attachment2.String != "" {
		attachments = append(attachments, data.Attachment2.String)
	}
	if data.Attachment3.Valid && data.Attachment3.String != "" {
		attachments = append(attachments, data.Attachment3.String)
	}
	if data.Attachment4.Valid && data.Attachment4.String != "" {
		attachments = append(attachments, data.Attachment4.String)
	}

	return domain.Mailbox{
		Id:          data.Id,
		Token:       data.Token.String,
		Status:      data.Status,
		From:        data.EmailFrom,
		Cc:          data.EmailCc,
		Subject:     data.EmailSubject,
		Body:        data.EmailBody,
		Attachments: attachments,
		Date:        data.EmailSendDate.Time,
		ParentId:    checkIfNull(data.ParentId),
		CreatedAt:   data.CreatedAt.Time,
		UpdatedAt:   data.UpdatedAt.Time,
	}
}

func checkIfNull(data sql.NullInt64) int64 {
	if data.Valid {
		return data.Int64
	}
	return 0
}

func MailboxDomainToEntityAdapter(data domain.Mailbox) entity.InbasketMailboxMySql {
	var attachment1 string
	var attachment2 string
	var attachment3 string
	var attachment4 string
	if len(data.Attachments) >= 4 {
		attachment1 = data.Attachments[0]
		attachment2 = data.Attachments[1]
		attachment3 = data.Attachments[2]
		attachment4 = data.Attachments[3]
	}
	if len(data.Attachments) == 3 {
		attachment1 = data.Attachments[0]
		attachment2 = data.Attachments[1]
		attachment3 = data.Attachments[2]
	}
	if len(data.Attachments) == 2 {
		attachment1 = data.Attachments[0]
		attachment2 = data.Attachments[1]
	}
	if len(data.Attachments) == 1 {
		attachment1 = data.Attachments[0]
	}
	return entity.InbasketMailboxMySql{
		Id:            data.Id,
		Token:         sql.NullString{String: data.Token, Valid: true},
		EmailFrom:     data.From,
		EmailCc:       data.Cc,
		EmailSubject:  data.Subject,
		EmailBody:     data.Body,
		EmailSendDate: sql.NullTime{Time: data.Date, Valid: true},
		Attachment1:   sql.NullString{String: attachment1, Valid: isValid(attachment1)},
		Attachment2:   sql.NullString{String: attachment2, Valid: isValid(attachment2)},
		Attachment3:   sql.NullString{String: attachment3, Valid: isValid(attachment3)},
		Attachment4:   sql.NullString{String: attachment4, Valid: isValid(attachment4)},
		Status:        data.Status,
		ParentId:      sql.NullInt64{Int64: data.ParentId, Valid: true},
		CreatedAt:     sql.NullTime{Time: data.CreatedAt, Valid: true},
		UpdatedAt:     sql.NullTime{Time: data.UpdatedAt, Valid: true},
	}
}
