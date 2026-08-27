package util

import (
	"tandur.com/src/adapter/out/mysql/one_on_one/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(data entity.OneOnOneMySql) domain.OneOnOne {
	var finalAssessors []domain.Assessor = []domain.Assessor{}
	finalAssessors = append(finalAssessors, domain.Assessor{
		Id: data.Assessor1Id,
	})
	if data.Assessor2Id.Valid {
		finalAssessors = append(finalAssessors, domain.Assessor{
			Id: data.Assessor2Id.Int64,
		})
	}
	if data.Assessor3Id.Valid {
		finalAssessors = append(finalAssessors, domain.Assessor{
			Id: data.Assessor3Id.Int64,
		})
	}
	if data.Assessor4Id.Valid {
		finalAssessors = append(finalAssessors, domain.Assessor{
			Id: data.Assessor4Id.Int64,
		})
	}
	return domain.OneOnOne{
		Token:      data.Token.String,
		MeetingUrl: data.MeetingUrl.String,
		StartDate:  data.StartDate.Time,
		EndDate:    data.EndDate.Time,
		Type:       data.ToolsType,
		Assessors:  finalAssessors,
	}
}
