package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/roleplay/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type RoleplayRepository struct {
	tableName string
}

func NewRoleplayRepository() RoleplayBaseRepository {
	return &RoleplayRepository{
		tableName: "V_GET_ROLEPLAY",
	}
}

func (repo *RoleplayRepository) GetAll(filter domain.RolePlay) (*[]entity.RoleplayMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, participant_instruction, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id, asesi, role_player_id FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.RoleplayMySql
	for results.Next() {
		var res entity.RoleplayMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.ParticipantInstruction, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id, &res.Asesi, &res.RolePlayerId)
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

func (repo *RoleplayRepository) GetById(id int64) (*entity.RoleplayMySql, error) {
	db := util.GetMySQL()
	var res entity.RoleplayMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, meeting_url, participant_instruction, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id, asesi, role_player_id FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.MeetingUrl, &res.ParticipantInstruction, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id, &res.Asesi, &res.RolePlayerId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *RoleplayRepository) GetByToken(token string) (*[]entity.RoleplayMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, tools_type, instruction_file, instruction, meeting_url, participant_instruction, start_date, end_date, assessor_1_id, assessor_2_id, assessor_3_id, assessor_4_id, asesi, role_player_id FROM %s WHERE token = ?", repo.tableName)
	log.Println(query)
	results, err := db.Query(query, token)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.RoleplayMySql
	for results.Next() {
		var res entity.RoleplayMySql
		err = results.Scan(&res.Token, &res.ToolsType, &res.InstructionFile, &res.Instruction, &res.MeetingUrl, &res.ParticipantInstruction, &res.StartDate, &res.EndDate, &res.Assessor1Id, &res.Assessor2Id, &res.Assessor3Id, &res.Assessor4Id, &res.Asesi, &res.RolePlayerId)
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
