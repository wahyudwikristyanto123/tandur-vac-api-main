package repository

import (
	"database/sql"
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/subtes/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type SubtesRepository struct {
	tableName string
}

func NewSubtesRepository() SubtesBaseRepository {
	return &SubtesRepository{
		tableName: "V_GET_SUBTES",
	}
}

func (repo *SubtesRepository) GetAll(filter domain.Subtes) (*[]entity.SubtesMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, id, user_name, assessment_center_id, title, start_date, end_date, tools_type, order_schedule, url_proctoring, status, duration FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.SubtesMySql
	for results.Next() {
		var res entity.SubtesMySql
		err = results.Scan(&res.Token, &res.ID, &res.UserName, &res.AssessmentCenterId, &res.Title, &res.StartDate, &res.EndDate, &res.ToolsType, &res.OrderSchedule, &res.URLProctoring, &res.Status, &res.Duration)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *SubtesRepository) GetById(id int64) (*entity.SubtesMySql, error) {
	db := util.GetMySQL()
	var res entity.SubtesMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, id, user_name, assessment_center_id, title, start_date, end_date, tools_type, type, order_schedule, url_proctoring, status, duration FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.ID, &res.UserName, &res.AssessmentCenterId, &res.Title, &res.StartDate, &res.EndDate, &res.ToolsType, &res.Type, &res.OrderSchedule, &res.URLProctoring, &res.Status, &res.Duration)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *SubtesRepository) GetByToken(token string) (*[]entity.SubtesMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, id, user_name, assessment_center_id, title, start_date, end_date, tools_type, type, order_schedule, url_proctoring, status, duration FROM %s WHERE token = '%s'", repo.tableName, token)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.SubtesMySql
	for results.Next() {
		var res entity.SubtesMySql
		err = results.Scan(&res.Token, &res.ID, &res.UserName, &res.AssessmentCenterId, &res.Title, &res.StartDate, &res.EndDate, &res.ToolsType, &res.Type, &res.OrderSchedule, &res.URLProctoring, &res.Status, &res.Duration)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *SubtesRepository) UpdateStatusById(id int64, status string) error {
	db := util.GetMySQL()
	query := fmt.Sprintf(`UPDATE trx_ap_class_management_detail SET status_test='%s' WHERE id=%d`, status, id)
	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SubtesRepository) UpdateResultById(id int64, result string) error {
	db := util.GetMySQL()
	query := fmt.Sprintf(`UPDATE trx_ap_class_management_detail SET result='%s', updated_at=NOW() WHERE id=%d`, result, id)
	log.Println(query)
	_, err := db.Query(query)
	if err != nil {
		return err
	}
	return nil
}

func (repo *SubtesRepository) GetResultById(id int64) (string, error) {
	db := util.GetMySQL()
	var res sql.NullString
	err := db.
		QueryRow(fmt.Sprintf("SELECT result FROM trx_ap_class_management_detail WHERE id = ?"), id).
		Scan(&res)
	if err != nil {
		return "", err
	}
	return res.String, nil
}
