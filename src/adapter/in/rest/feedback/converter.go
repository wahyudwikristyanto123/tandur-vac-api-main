package feedback

import "tandur.com/src/domain"

func RequestToDomain(data FeedbackRequest) domain.Feedback {
	return domain.Feedback{
		Token:              data.Token,
		AssessmentCenterId: data.AcId,
		Feedback:           data.Feedback,
	}
}
