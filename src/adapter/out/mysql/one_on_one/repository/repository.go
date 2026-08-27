package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/one_on_one/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type OneOnOneRepository struct {
	tableName string
}

func NewOneOnOneRepository() OneOnOneBaseRepository {
	return &OneOnOneRepository{
		tableName: "V_GET_1ON1",
	}
}

func (repo *OneOnOneRepository) GetAll(filter domain.OneOnOne) (*[]entity.OneOnOneMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, meeting_url, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.OneOnOneMySql
	for results.Next() {
		var res entity.OneOnOneMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.MeetingUrl, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *OneOnOneRepository) GetById(id int64) (*entity.OneOnOneMySql, error) {
	db := util.GetMySQL()
	var res entity.OneOnOneMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, tools_type, meeting_url, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.ToolsType, &res.MeetingUrl, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *OneOnOneRepository) GetByToken(token string) (*[]entity.OneOnOneMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, meeting_url, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id FROM %s WHERE token = '%s'", repo.tableName, token)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.OneOnOneMySql
	for results.Next() {
		var res entity.OneOnOneMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.MeetingUrl, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}
