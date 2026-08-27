package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/feedback/entity"
	"tandur.com/src/util"
)

type FeedbackRepository struct {
	tableName string
}

func NewFeedbackRepository() FeedbackBaseRepository {
	return &FeedbackRepository{
		tableName: "trx_ac_feedback",
	}
}

func (repo *FeedbackRepository) GetFeedbackByToken(token string) (*[]entity.FeedbackMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, id, trx_assessment_center_id, feedback, created_at, updated_at FROM %s WHERE token = ? ORDER BY updated_at DESC", repo.tableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.FeedbackMySql
	for results.Next() {
		var res entity.FeedbackMySql
		err = results.Scan(&res.Token, &res.Id, &res.AssessmentCenterId, &res.Feedback, &res.CreatedAt, &res.UpdatedAt)
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

func (repo *FeedbackRepository) GetFeedbackById(id int64) (*entity.FeedbackMySql, error) {
	db := util.GetMySQL()
	var res entity.FeedbackMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, id, trx_assessment_center_id, feedback, created_at, updated_at FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.Id, &res.AssessmentCenterId, &res.Feedback, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *FeedbackRepository) InsertFeedback(data entity.FeedbackMySql) (int64, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf(`INSERT INTO %s 
	 (token, trx_assessment_center_id, feedback, created_at, updated_at) 
	 VALUES (?, ?, ?, NOW(), NOW())`,
		repo.tableName,
	)
	result, err := db.Exec(query, data.Token.String, data.AssessmentCenterId, data.Feedback.String)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (repo *FeedbackRepository) UpdateFeedback(data entity.FeedbackMySql) error {
	db := util.GetMySQL()
	query := fmt.Sprintf(`UPDATE %s SET token=?, trx_assessment_center_id=?, feedback=?, updated_at=NOW() WHERE id=?`,
		repo.tableName,
	)
	_, err := db.Exec(query, data.Token.String, data.AssessmentCenterId, data.Feedback.String, data.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FeedbackRepository) DeleteFeedback(id int64) error {
	db := util.GetMySQL()
	query := fmt.Sprintf("DELETE FROM %s WHERE id=?", repo.tableName)
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
