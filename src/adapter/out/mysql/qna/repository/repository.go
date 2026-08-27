package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/qna/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type QnaRepository struct {
	tableName string
}

func NewQnaRepository() QnaBaseRepository {
	return &QnaRepository{
		tableName: "V_GET_QNA_MASTER",
	}
}

func (repo *QnaRepository) GetAll(filter domain.Qna) (*[]entity.QnaMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT qna_id, detail_id, token, instruction_file, instruction, type_qna, start_date, end_date FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.QnaMySql
	for results.Next() {
		var res entity.QnaMySql
		err = results.Scan(&res.Id, &res.IdQna, &res.Token, &res.InstructionFile, &res.Instruction, &res.TypeQna, &res.StartDate, &res.EndDate)
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

func (repo *QnaRepository) GetDetailsById(id int64) (*[]entity.QnaDetailMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT id, trx_ac_question_answer_id, head_question, question, created_at, updated_at FROM %s WHERE trx_ac_question_answer_id = ?", "trx_ac_question_answer_detail")
	log.Println(query)
	results, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.QnaDetailMySql
	for results.Next() {
		var res entity.QnaDetailMySql
		err = results.Scan(&res.Id, &res.TrxAcQuestionAnswerId, &res.HeadQuestion, &res.Question, &res.CreatedAt, &res.UpdatedAt)
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

func (repo *QnaRepository) GetById(id int64) (*entity.QnaMySql, error) {
	db := util.GetMySQL()
	var res entity.QnaMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT qna_id, detail_id, token, instruction_file, instruction, type_qna, start_date, end_date FROM %s WHERE detail_id = ?", repo.tableName), id).
		Scan(&res.Id, &res.IdQna, &res.Token, &res.InstructionFile, &res.Instruction, &res.TypeQna, &res.StartDate, &res.EndDate)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *QnaRepository) GetByToken(token string) (*[]entity.QnaMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT qna_id, detail_id, token, instruction_file, instruction, type_qna, start_date, end_date FROM %s WHERE token = ?", repo.tableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.QnaMySql
	for results.Next() {
		var res entity.QnaMySql
		err = results.Scan(&res.Id, &res.IdQna, &res.Token, &res.InstructionFile, &res.Instruction, &res.TypeQna, &res.StartDate, &res.EndDate)
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

func (repo *QnaRepository) UpsertResult(data entity.QnaResultRequestMysql) error {
	db := util.GetMySQL()
	query := `INSERT INTO trx_ac_qna_result 
			(question_id, qna_id, token, result) 
			VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE result = VALUES(result)`
	_, err := db.Exec(query, data.QuestionId, data.QnaId, data.Token, data.Result.String)
	if err != nil {
		return err
	}
	return nil
}

func (repo *QnaRepository) GetResultsByToken(token string) (*[]entity.QnaResultRequestMysql, error) {
	db := util.GetMySQL()
	query := "SELECT question_id, qna_id, token, result FROM trx_ac_qna_result WHERE token = ?"
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.QnaResultRequestMysql
	for results.Next() {
		var res entity.QnaResultRequestMysql
		err = results.Scan(&res.QuestionId, &res.QnaId, &res.Token, &res.Result)
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
