package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/assessor/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type AssessorRepository struct {
	tableName string
}

func NewAssessorRepository() AssessorBaseRepository {
	return &AssessorRepository{
		tableName: "V_GET_AP_USER",
	}
}

func (repo *AssessorRepository) GetAll(filter domain.Assessor) (*[]entity.ApUserMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT id, user_code, user_name, user_type, name, email, photo, phone_number FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var data []entity.ApUserMySql
	for results.Next() {
		var res entity.ApUserMySql
		err = results.Scan(&res.Id, &res.UserCode, &res.UserName, &res.UserType, &res.Name, &res.Email, &res.Photo, &res.PhoneNumber)
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

func (repo *AssessorRepository) GetById(id int64) (*entity.ApUserMySql, error) {
	db := util.GetMySQL()
	var res entity.ApUserMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT id, user_code, user_name, user_type, name, email, photo, phone_number FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Id, &res.UserCode, &res.UserName, &res.UserType, &res.Name, &res.Email, &res.Photo, &res.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
