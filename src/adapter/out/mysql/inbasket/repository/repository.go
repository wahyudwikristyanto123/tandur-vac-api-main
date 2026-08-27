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
	query := fmt.Sprintf("SELECT token, tools_type, event_id, event_title, event_date, event_description, start_date, end_date FROM %s WHERE token = ?", repo.eventTableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.InbasketEventMySql
	for results.Next() {
		var res entity.InbasketEventMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.EventId, &res.EventTitle, &res.EventDate, &res.EventDescription, &res.StartDate, &res.EndDate)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *InbasketRepository) GetEmailsByToken(token string) (*[]entity.InbasketEmailMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, email_id, email_from, email_cc, email_subject, email_body, email_send_date, created_at, updated_at, attachment_1, attachment_2, attachment_3, attachment_4, start_date, end_date FROM %s WHERE token = ?", repo.emailTableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.InbasketEmailMySql
	for results.Next() {
		var res entity.InbasketEmailMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.EmailId, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.CreatedAt, &res.UpdatedAt, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.StartDate, &res.EndDate)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *InbasketRepository) GetMailboxByToken(token string, statuses []string) (*[]entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	if len(statuses) == 0 {
		statuses = []string{"UNREAD", "READ", "DRAFT", "SENT", "REPLIED"}
	}
	placeholders := make([]string, len(statuses))
	args := make([]interface{}, 0, len(statuses)+1)
	args = append(args, token)
	for i, s := range statuses {
		placeholders[i] = "?"
		args = append(args, s)
	}
	query := fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE token = ? AND status IN (%s) ORDER BY updated_at DESC", repo.mailboxTableName, strings.Join(placeholders, ", "))
	log.Println(query)
	results, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.InbasketMailboxMySql
	for results.Next() {
		var res entity.InbasketMailboxMySql
		err = results.Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
	}
	return &data, nil
}

func (repo *InbasketRepository) SearchMailboxByToken(q string, token string, statuses []string) (*[]entity.InbasketMailboxMySql, error) {
	db := util.GetMySQL()
	if len(statuses) == 0 {
		statuses = []string{"UNREAD", "READ", "DRAFT", "SENT", "REPLIED"}
	}
	placeholders := make([]string, len(statuses))
	args := make([]interface{}, 0, len(statuses)+3)
	args = append(args, token)
	for i, s := range statuses {
		placeholders[i] = "?"
		args = append(args, s)
	}
	likePattern := "%" + q + "%"
	args = append(args, likePattern, likePattern)

	query := fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE token = ? AND status IN (%s) AND (email_subject LIKE ? OR email_body LIKE ?) ORDER BY updated_at DESC", repo.mailboxTableName, strings.Join(placeholders, ", "))
	log.Println(query)
	results, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.InbasketMailboxMySql
	for results.Next() {
		var res entity.InbasketMailboxMySql
		err = results.Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
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
	query := fmt.Sprintf("SELECT id, email_from, email_cc, email_subject, email_body, email_send_date, attachment_1, attachment_2, attachment_3, attachment_4, created_at, updated_at, token, status, parent_id FROM %s WHERE parent_id = ? ORDER BY updated_at DESC", repo.mailboxTableName)
	log.Println(query)
	results, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.InbasketMailboxMySql
	for results.Next() {
		var res entity.InbasketMailboxMySql
		err = results.Scan(&res.Id, &res.EmailFrom, &res.EmailCc, &res.EmailSubject, &res.EmailBody, &res.EmailSendDate, &res.Attachment1, &res.Attachment2, &res.Attachment3, &res.Attachment4, &res.CreatedAt, &res.UpdatedAt, &res.Token, &res.Status, &res.ParentId)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	if err = results.Err(); err != nil {
		return nil, err
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
	 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), ?, ?, ?)`,
		repo.mailboxTableName,
	)
	var sendDateStr string
	if !data.EmailSendDate.Time.IsZero() {
		sendDateStr = data.EmailSendDate.Time.Format("2006-01-02 15:04:05")
	}
	result, err := db.Exec(query,
		data.EmailFrom, data.EmailCc, data.EmailSubject, data.EmailBody, sendDateStr,
		data.Attachment1.String, data.Attachment2.String, data.Attachment3.String, data.Attachment4.String,
		data.Token.String, data.Status, data.ParentId.Int64,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *InbasketRepository) InsertMultiMailbox(data []entity.InbasketMailboxMySql) error {
	if len(data) == 0 {
		return nil
	}
	db := util.GetMySQL()
	var valuePlaceholders []string
	var args []interface{}
	for _, item := range data {
		valuePlaceholders = append(valuePlaceholders, "(?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW(), ?, ?, ?)")
		var sendDateStr string
		if !item.EmailSendDate.Time.IsZero() {
			sendDateStr = item.EmailSendDate.Time.Format("2006-01-02 15:04:05")
		}
		args = append(args,
			item.EmailFrom, item.EmailCc, item.EmailSubject, item.EmailBody, sendDateStr,
			item.Attachment1.String, item.Attachment2.String, item.Attachment3.String, item.Attachment4.String,
			item.Token.String, item.Status, item.ParentId.Int64,
		)
	}
	query := fmt.Sprintf(`INSERT INTO %s 
			(email_from, email_cc, email_subject, email_body, email_send_date, 
				attachment_1, attachment_2, attachment_3, attachment_4, 
				created_at, updated_at, token, status, parent_id) 
			VALUES %s`, repo.mailboxTableName, strings.Join(valuePlaceholders, ", "))

	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InbasketRepository) UpdateMailbox(data entity.InbasketMailboxMySql) error {
	db := util.GetMySQL()
	query := fmt.Sprintf(`UPDATE %s SET email_from=?, email_cc=?, email_subject=?, email_body=?, email_send_date=?, 
	attachment_1=?, attachment_2=?, attachment_3=?, attachment_4=?, 
	updated_at=NOW(), token=?, status=?, parent_id=? WHERE id=?`,
		repo.mailboxTableName,
	)
	var sendDateStr string
	if !data.EmailSendDate.Time.IsZero() {
		sendDateStr = data.EmailSendDate.Time.Format("2006-01-02 15:04:05")
	}
	_, err := db.Exec(query,
		data.EmailFrom, data.EmailCc, data.EmailSubject, data.EmailBody, sendDateStr,
		data.Attachment1.String, data.Attachment2.String, data.Attachment3.String, data.Attachment4.String,
		data.Token.String, data.Status, data.ParentId.Int64, data.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InbasketRepository) DeleteMailbox(id int64) error {
	db := util.GetMySQL()
	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", repo.mailboxTableName)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *InbasketRepository) DeleteByParentIdStatus(id int64, status string) error {
	db := util.GetMySQL()
	query := fmt.Sprintf("DELETE FROM %s WHERE parent_id=? AND status = ?", repo.mailboxTableName)
	_, err := db.Exec(query, id, status)
	if err != nil {
		return err
	}
	return nil
}
