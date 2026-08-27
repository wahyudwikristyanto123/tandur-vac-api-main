package usecase

import "tandur.com/src/domain"

type FeedbackUseCase interface {
	SubmitFeedback(data domain.Feedback) error
}
