package repository

import (
	"fmt"
	"log"
	"strings"

	"tandur.com/src/adapter/out/mysql/inbasket/entity"
	"tandur.com/src/util"
)

type InbasketRepository struct {
	mainTableName    string
	eventTableName   string
	emailTableName   string
	mailboxTableName string
}

func NewInbasketRepository() InbasketBaseRepository {
	return &InbasketRepository{
		mainTableName:    "V_GET_INBASKET_FILE",
		eventTableName:   "V_GET_INBASKET_EVENT",
		emailTableName:   "V_GET_INBASKET_EMAIL",
		mailboxTableName: "trx_ac_inbasket_mailbox",
	}
}

func (repo *InbasketRepository) GetFileByToken(token string) (*entity.InbasketFileMySql, error) {
	db := util.GetMySQL()
	var res entity.InbasketFileMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, tools_type, instruction, instruction_file, company_profile_file, organization_structure_file, start_date, end_date FROM %s WHERE token = ?", repo.mainTableName), token).
		Scan(&res.Token, &res.ToolsType, &res.Instruction, &res.InstructionFile, &res.CompanyProfileFile, &res.OrganizationStructureFile, &res.StartDate, &res.EndDate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *InbasketRepository) GetEventsByToken(token string) (*[]entity.InbasketEventMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, event_id, event_title, event_date, event_description, start_date, end_date FROM %s WHERE token = '%s'", repo.eventTableName, token)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.InbasketEventMySql
	for results.Next() {
		var res entity.InbasketEventMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.EventId, &res.EventTitle, &res.EventDate, &res.EventDescription, &res.StartDate, &res.EndDate)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *InbasketRepository) GetEmailsByToken(token string) (*[]entity.InbasketEmailMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, email_id, email_from, email_cc, email_subject, email_body, email_send_date, created_at, updated_at, attachment_1, attachment_2, attachment_3, attachment_4, start_date, end_date FROM %s WHERE token = '%s'", repo.emailTableName, token)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.InbasketEmailMySql
	for results.Next() {
		var res entity.InbasketEmailMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.EmailId, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.CreatedAt, &res.UpdatedAt, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.StartDate, &res.EndDate)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *InbasketRepository) GetMailboxByToken(token string, statuses []string) (*[]entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	var statusesWithQuotes []string = []string{}
	for i := 0; i < len(statuses); i++ {
		statusesWithQuotes = append(statusesWithQuotes, fmt.Sprintf("'%s'", statuses[i]))
	}
	query := fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE token = '%s' AND status IN (%s) ORDER BY updated_at DESC", repo.mailboxTableName, token, strings.Join(statusesWithQuotes, ", "))
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.InbasketMailboxMySql
	for results.Next() {
		var res entity.InbasketMailboxMySql
		err = results.Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *InbasketRepository) SearchMailboxByToken(q string, token string, statuses []string) (*[]entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	var statusesWithQuotes []string = []string{}
	for i := 0; i < len(statuses); i++ {
		statusesWithQuotes = append(statusesWithQuotes, fmt.Sprintf("'%s'", statuses[i]))
	}
	query := fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE token = '%s' AND status IN (%s) AND (email_subject OR email_body LIKE '%%%s%%') ORDER BY updated_at DESC", repo.mailboxTableName, token, strings.Join(statusesWithQuotes, ", "), q)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.InbasketMailboxMySql
	for results.Next() {
		var res entity.InbasketMailboxMySql
		err = results.Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *InbasketRepository) GetMailboxById(id int64) (*entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	var res entity.InbasketMailboxMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE id = ?", repo.mailboxTableName), id).
		Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *InbasketRepository) GetMailboxByParentId(id int64) (*[]entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE parent_id = %d ORDER BY updated_at DESC", repo.mailboxTableName, id)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.InbasketMailboxMySql
	for results.Next() {
		var res entity.InbasketMailboxMySql
		err = results.Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *InbasketRepository) GetMailboxByParentIdStatus(parentId int64, status string) (*entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	var res entity.InbasketMailboxMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE parent_id = ? AND status = ?", repo.mailboxTableName), parentId, status).
		Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *InbasketRepository) InsertMailbox(data entity.InbasketMailboxMySql) (int64, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf(`INSERT INTO %s 
	 (email_from, email_cc, email_subject, email_body, email_send_date, 
		attachment_1, attachment_2, attachment_3, attachment_4, 
		created_at, updated_at, token, status, parent_id) 
	 VALUES ('%s', '%s', '%s', '%s', '%s',
	 '%s', '%s', '%s', '%s',
		NOW(), NOW(), '%s', '%s', %d)`,
		repo.mailboxTableName, data.EmailFrom, data.EmailCc, data.EmailSubject, data.EmailBody, data.EmailSendDate.Time.Format("2006-01-02 15:04:05"),
		data.Attachment1.String, data.Attachment2.String, data.Attachment3.String, data.Attachment4.String,
		data.Token.String, data.Status, data.ParentId.Int64)
	// log.Println(query)
	result, err := db.Exec(query)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *InbasketRepository) InsertMultiMailbox(data []entity.InbasketMailboxMySql) error {
	db := util.GetMySQL()
	var query string = fmt.Sprintf(`INSERT INTO %s 
			(email_from, email_cc, email_subject, email_body, email_send_date, 
				attachment_1, attachment_2, attachment_3, attachment_4, 
				created_at, updated_at, token, status, parent_id) 
			VALUES `, repo.mailboxTableName)
	var valuePlaceholders []string
	for i := 0; i < len(data); i++ {
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf(`('%s', '%s', '%s', '%s', '%s',
			'%s', '%s', '%s', '%s',
				NOW(), NOW(), '%s', '%s', %d)`,
			data[i].EmailFrom, data[i].EmailCc, data[i].EmailSubject, strings.ReplaceAll(data[i].EmailBody, "'", "\\'"), data[i].EmailSendDate.Time.Format("2006-01-02 15:04:05"),
			data[i].Attachment1.String, data[i].Attachment2.String, data[i].Attachment3.String, data[i].Attachment4.String,
			data[i].Token.String, data[i].Status, data[i].ParentId.Int64))
	}
	query += strings.Join(valuePlaceholders, ", ")
	// log.Println(query)
	// logger.WARN(query)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InbasketRepository) UpdateMailbox(data entity.InbasketMailboxMySql) error {
	db := util.GetMySQL()
	query := fmt.Sprintf(`UPDATE %s SET email_from='%s', email_cc='%s', email_subject='%s', email_body='%s', email_send_date='%s', 
	attachment_1='%s', attachment_2='%s', attachment_3='%s', attachment_4='%s',
	updated_at=NOW(), token='%s', status='%s', parent_id=%d WHERE id=%d`,
		repo.mailboxTableName, data.EmailFrom, data.EmailCc, data.EmailSubject, data.EmailBody, data.EmailSendDate.Time.Format("2006-01-02 15:04:05"),
		data.Attachment1.String, data.Attachment2.String, data.Attachment3.String, data.Attachment4.String,
		data.Token.String, data.Status, data.ParentId.Int64, data.Id)
	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InbasketRepository) DeleteMailbox(id int64) error {
	db := util.GetMySQL()
	query := fmt.Sprintf("DELETE FROM %s WHERE id=%d", repo.mailboxTableName, id)
	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InbasketRepository) DeleteByParentIdStatus(id int64, status string) error {
	db := util.GetMySQL()
	query := fmt.Sprintf("DELETE FROM %s WHERE parent_id=%d AND status = '%s'", repo.mailboxTableName, id, status)
	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		return err
	}
	return nil
}
