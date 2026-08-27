package service

import (
	"tandur.com/src/adapter/out/mysql/feedback/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type FeedbackService struct {
	adapter *adapter.FeedbackBaseAdapter
}

func NewFeedbackService(adapter *adapter.FeedbackBaseAdapter) usecase.FeedbackUseCase {
	return &FeedbackService{
		adapter: adapter,
	}
}

func (service *FeedbackService) SubmitFeedback(data domain.Feedback) error {
	results, err := (*service.adapter).GetFeedbackByToken(data.Token)
	if err != nil {
		return err
	}
	finalData := domain.Feedback{
		Token:              data.Token,
		AssessmentCenterId: data.AssessmentCenterId,
		Feedback:           data.Feedback,
	}
	if len(*results) < 1 {
		_, err = (*service.adapter).AddFeedback(finalData)
		if err != nil {
			return err
		}
	} else {
		finalData.Id = (*results)[0].Id
		err = (*service.adapter).UpdateFeedback(finalData)
		if err != nil {
			return err
		}
	}

	return nil
}
