package util

import (
	"strconv"
	"strings"

	"tandur.com/src/adapter/out/mysql/roleplay/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(data entity.RoleplayMySql) domain.RolePlay {
	var finalAssessors []domain.Assessor = []domain.Assessor{}
	var finalParticipants []domain.Assessor = []domain.Assessor{}
	var finalRoleplayers []domain.Assessor = []domain.Assessor{}
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
	if data.Asesi.Valid {
		participants := strings.Split(data.Asesi.String, ",")
		for i := 0; i < len(participants); i++ {
			id, err := strconv.ParseInt(participants[i], 10, 64)
			if err != nil {
				continue
			}
			finalParticipants = append(finalParticipants, domain.Assessor{
				Id: id,
			})
		}
	}
	if data.RolePlayerId.Valid {
		finalRoleplayers = append(finalRoleplayers, domain.Assessor{
			Id: data.RolePlayerId.Int64,
		})
	}
	return domain.RolePlay{
		Token:        data.Token.String,
		MeetingUrl:   data.MeetingUrl.String,
		StartDate:    data.StartDate.Time,
		EndDate:      data.EndDate.Time,
		Type:         data.ToolsType,
		Assessors:    finalAssessors,
		Participants: finalParticipants,
		Roleplayers:  finalRoleplayers,
	}
}
