package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/lgd/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type LgdRepository struct {
	tableName string
}

func NewLgdRepository() LgdBaseRepository {
	return &LgdRepository{
		tableName: "V_GET_LGD",
	}
}

func (repo *LgdRepository) GetAll(filter domain.Lgd) (*[]entity.LgdMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, participant_instruction, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id, asesi FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.LgdMySql
	for results.Next() {
		var res entity.LgdMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.ParticipantInstruction, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id, &res.Asesi)
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

func (repo *LgdRepository) GetById(id int64) (*entity.LgdMySql, error) {
	db := util.GetMySQL()
	var res entity.LgdMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, meeting_url, participant_instruction, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id, asesi FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.MeetingUrl, &res.ParticipantInstruction, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id, &res.Asesi)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *LgdRepository) GetByToken(token string) (*[]entity.LgdMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, meeting_url, participant_instruction, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id, asesi FROM %s WHERE token = ?", repo.tableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.LgdMySql
	for results.Next() {
		var res entity.LgdMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.MeetingUrl, &res.ParticipantInstruction, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id, &res.Asesi)
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
