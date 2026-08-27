package adapter

import (
	"tandur.com/src/domain"
)

type FeedbackBaseAdapter interface {
	GetFeedbackByToken(token string) (*[]domain.Feedback, error)
	GetFeedbackById(id int64) (*domain.Feedback, error)
	AddFeedback(data domain.Feedback) (int64, error)
	UpdateFeedback(data domain.Feedback) error
	RemoveFeedback(id int64) error
}
