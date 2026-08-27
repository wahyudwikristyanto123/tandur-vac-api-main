package repository

import (
	"tandur.com/src/adapter/out/mysql/feedback/entity"
)

type FeedbackBaseRepository interface {
	GetFeedbackByToken(token string) (*[]entity.FeedbackMySql, error)
	GetFeedbackById(id int64) (*entity.FeedbackMySql, error)
	InsertFeedback(data entity.FeedbackMySql) (int64, error)
	UpdateFeedback(data entity.FeedbackMySql) error
	DeleteFeedback(id int64) error
}
