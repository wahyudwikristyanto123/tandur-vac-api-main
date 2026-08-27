package util

import (
	"tandur.com/src/adapter/out/mysql/assessor/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(data entity.ApUserMySql) domain.Assessor {
	return domain.Assessor{
		Id:       data.Id,
		Username: data.UserName.String,
		Name:     data.Name,
		Email:    data.Email,
		ImageUrl: data.Photo,
		Role:     data.UserType,
	}
}
