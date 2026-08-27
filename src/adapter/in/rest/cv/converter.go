package cv

import "tandur.com/src/domain"

func RequestToDomain(data SubmitRequest) domain.Cv {
	return domain.Cv{
		Token:              data.Token,
		CmdId:              data.CmdId,
		AssessmentCenterId: data.AcId,
		Page1:              data.Page1,
		Page2:              data.Page2,
		Page3:              data.Page3,
		Page4:              data.Page4,
		Page5:              data.Page5,
		Page6:              data.Page6,
		Page7:              data.Page7,
		Page8:              data.Page8,
		Page9:              data.Page9,
		Page10:             data.Page10,
		Page11:             data.Page11,
		Page12:             data.Page12,
		Page13:             data.Page13,
	}
}
