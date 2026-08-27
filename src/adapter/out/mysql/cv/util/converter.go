package util

import (
	"database/sql"

	"tandur.com/src/adapter/out/mysql/cv/entity"
	"tandur.com/src/domain"
)

func isValid(val string) bool {
	return val != ""
}

func EntityAdapterToDomain(data entity.CvViewMySql) domain.Cv {
	return domain.Cv{
		Id:                 data.Id,
		Token:              data.Token.String,
		CmdId:              data.CmdId,
		AssessmentCenterId: data.AssessmentCenterId,
		Page1:              data.Page1.String,
		Page2:              data.Page2.String,
		Page3:              data.Page3.String,
		Page4:              data.Page4.String,
		Page5:              data.Page5.String,
		Page6:              data.Page6.String,
		Page7:              data.Page7.String,
		Page8:              data.Page8.String,
		Page9:              data.Page9.String,
		Page10:             data.Page10.String,
		Page11:             data.Page11.String,
		Page12:             data.Page12.String,
		Page13:             data.Page13.String,
		CreatedAt:          data.CreatedAt.Time,
		UpdatedAt:          data.UpdatedAt.Time,
	}
}

func DomainToEntityAdapter(data domain.Cv) entity.CvMySql {
	return entity.CvMySql{
		Id:                 data.Id,
		Token:              sql.NullString{String: data.Token, Valid: true},
		AssessmentCenterId: data.AssessmentCenterId,
		Page1:              sql.NullString{String: data.Page1, Valid: true},
		Page2:              sql.NullString{String: data.Page2, Valid: true},
		Page3:              sql.NullString{String: data.Page3, Valid: true},
		Page4:              sql.NullString{String: data.Page4, Valid: true},
		Page5:              sql.NullString{String: data.Page5, Valid: true},
		Page6:              sql.NullString{String: data.Page6, Valid: true},
		Page7:              sql.NullString{String: data.Page7, Valid: true},
		Page8:              sql.NullString{String: data.Page8, Valid: true},
		Page9:              sql.NullString{String: data.Page9, Valid: true},
		Page10:             sql.NullString{String: data.Page10, Valid: true},
		Page11:             sql.NullString{String: data.Page11, Valid: true},
		Page12:             sql.NullString{String: data.Page12, Valid: true},
		Page13:             sql.NullString{String: data.Page13, Valid: true},
		CreatedAt:          sql.NullTime{Time: data.CreatedAt, Valid: true},
		UpdatedAt:          sql.NullTime{Time: data.UpdatedAt, Valid: true},
	}
}

func checkIfNull(data sql.NullInt64) int64 {
	if data.Valid {
		return data.Int64
	}
	return 0
}
