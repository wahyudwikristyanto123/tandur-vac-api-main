package util

import (
	"database/sql"

	"tandur.com/src/adapter/out/mysql/qna/entity"
	"tandur.com/src/domain"
)

func EntityAdapterToDomain(data entity.QnaMySql) domain.Qna {
	return domain.Qna{
		Id:             data.Id,
		IdQna:          data.IdQna,
		Token:          data.Token.String,
		Type:           data.TypeQna.String,
		InstructionUrl: data.InstructionFile.String,
		Instruction:    data.Instruction,
		StartDate:      data.StartDate.Time,
		EndDate:        data.EndDate.Time,
	}
}

func EntityAdapterDetailToDomainDetail(data entity.QnaDetailMySql) domain.QnaDetail {
	return domain.QnaDetail{
		Id:       data.Id,
		ParentId: data.TrxAcQuestionAnswerId,
		Title:    data.HeadQuestion,
		Question: data.Question,
	}
}

func DomainToEntityAdapterResult(data domain.QnaResultRequest) entity.QnaResultRequestMysql {
	return entity.QnaResultRequestMysql{
		QuestionId: data.QuestionId,
		QnaId:      data.QnaId,
		Token:      data.Token,
		Result:     sql.NullString{String: data.Result, Valid: true},
	}
}

func EntityAdapterResultToDomain(data entity.QnaResultRequestMysql) domain.QnaResultRequest {
	return domain.QnaResultRequest{
		QuestionId: data.QuestionId,
		QnaId:      data.QnaId,
		Token:      data.Token,
		Result:     data.Result.String,
	}
}
