package util

import (
	"database/sql"

	"tandur.com/src/adapter/out/mysql/feedback/entity"
	"tandur.com/src/domain"
)

func isValid(val string) bool {
	return val != ""
}

func EntityAdapterToDomain(data entity.FeedbackMySql) domain.Feedback {
	return domain.Feedback{
		Id:                 data.Id,
		Token:              data.Token.String,
		AssessmentCenterId: data.AssessmentCenterId,
		Feedback:           data.Feedback.String,
		CreatedAt:          data.CreatedAt.Time,
		UpdatedAt:          data.UpdatedAt.Time,
	}
}

func DomainToEntityAdapter(data domain.Feedback) entity.FeedbackMySql {
	return entity.FeedbackMySql{
		Id:                 data.Id,
		Token:              sql.NullString{String: data.Token, Valid: true},
		AssessmentCenterId: data.AssessmentCenterId,
		Feedback:           sql.NullString{String: data.Feedback, Valid: true},
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
