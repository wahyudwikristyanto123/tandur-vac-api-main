package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/cv/entity"
	"tandur.com/src/util"
)

type CvRepository struct {
	tableName string
	viewName  string
}

func NewCvRepository() CvBaseRepository {
	return &CvRepository{
		tableName: "trx_ac_cv",
		viewName:  "V_GET_CV",
	}
}

func (repo *CvRepository) GetCvByToken(token string) (*[]entity.CvViewMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, cmd_id, id, trx_assessment_center_id, page1, page2, page3, page4, page5, page6, page7, page8, page9, page10, page11, page12, page13, created_at, updated_at FROM %s WHERE token = ? ORDER BY updated_at DESC", repo.viewName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.CvViewMySql
	for results.Next() {
		var res entity.CvViewMySql
		err = results.Scan(&res.Token, &res.CmdId, &res.Id, &res.AssessmentCenterId, &res.Page1, &res.Page2, &res.Page3, &res.Page4, &res.Page5, &res.Page6, &res.Page7, &res.Page8, &res.Page9, &res.Page10, &res.Page11, &res.Page12, &res.Page13, &res.CreatedAt, &res.UpdatedAt)
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

func (repo *CvRepository) GetCvById(id int64) (*entity.CvViewMySql, error) {
	db := util.GetMySQL()
	var res entity.CvViewMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, cmd_id, id, trx_assessment_center_id, page1, page2, page3, page4, page5, page6, page7, page8, page9, page10, page11, page12, page13, created_at, updated_at FROM %s WHERE id = ?", repo.viewName), id).
		Scan(&res.Token, &res.CmdId, &res.Id, &res.AssessmentCenterId, &res.Page1, &res.Page2, &res.Page3, &res.Page4, &res.Page5, &res.Page6, &res.Page7, &res.Page8, &res.Page9, &res.Page10, &res.Page11, &res.Page12, &res.Page13, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *CvRepository) InsertCv(data entity.CvMySql) (int64, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf(`INSERT INTO %s 
	 (token, trx_assessment_center_id, page1, page2, page3, page4, 
		page5, page6, page7, page8, page9, page10, page11, page12, page13, 
		created_at, updated_at) 
	 VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`,
		repo.tableName,
	)
	result, err := db.Exec(query,
		data.Token.String, data.AssessmentCenterId, data.Page1.String, data.Page2.String, data.Page3.String, data.Page4.String,
		data.Page5.String, data.Page6.String, data.Page7.String, data.Page8.String, data.Page9.String, data.Page10.String,
		data.Page11.String, data.Page12.String, data.Page13.String,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *CvRepository) UpdateCv(data entity.CvMySql) error {
	db := util.GetMySQL()
	query := fmt.Sprintf(`UPDATE %s SET token=?, trx_assessment_center_id=?, page1=?, page2=?, page3=?, page4=?, 
	page5=?, page6=?, page7=?, page8=?, page9=?, page10=?, 
	page11=?, page12=?, page13=?, updated_at=NOW() WHERE id=?`,
		repo.tableName,
	)
	_, err := db.Exec(query,
		data.Token.String, data.AssessmentCenterId, data.Page1.String, data.Page2.String, data.Page3.String, data.Page4.String,
		data.Page5.String, data.Page6.String, data.Page7.String, data.Page8.String, data.Page9.String, data.Page10.String,
		data.Page11.String, data.Page12.String, data.Page13.String, data.Id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CvRepository) DeleteCv(id int64) error {
	db := util.GetMySQL()
	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", repo.tableName)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
