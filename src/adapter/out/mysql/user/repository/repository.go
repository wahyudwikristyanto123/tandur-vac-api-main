package repository

import (
	"fmt"
	"log"

	"tandur.com/src/adapter/out/mysql/user/entity"
	"tandur.com/src/domain"
	"tandur.com/src/util"
)

type UserRepository struct {
	tableName string
}

func NewUserRepository() UserBaseRepository {
	return &UserRepository{
		tableName: "V_GET_LOGIN",
	}
}

func (repo *UserRepository) GetAll(filter domain.User) (*[]entity.UserMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, username, email, password, asesi FROM %s", repo.tableName)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.UserMySql
	for results.Next() {
		var res entity.UserMySql
		err = results.Scan(&res.Token, &res.Username, &res.Email, &res.Password, &res.Asesi)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}

func (repo *UserRepository) GetById(id int64) (*entity.UserMySql, error) {
	db := util.GetMySQL()
	var res entity.UserMySql
	err := db.
		QueryRow(fmt.Sprintf("SELECT token, username, email, password, asesi FROM %s WHERE id = ?", repo.tableName), id).
		Scan(&res.Token, &res.Username, &res.Email, &res.Password, &res.Asesi)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (repo *UserRepository) GetByToken(token string) (*[]entity.UserMySql, error) {
	db := util.GetMySQL()
	query := fmt.Sprintf("SELECT token, username, email, password, asesi FROM %s WHERE token = '%s'", repo.tableName, token)
	log.Println(query)
	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var data []entity.UserMySql
	for results.Next() {
		var res entity.UserMySql
		err = results.Scan(&res.Token, &res.Username, &res.Email, &res.Password, &res.Asesi)
		if err != nil {
			return nil, err
		}
		data = append(data, res)
	}
	return &data, nil
}
