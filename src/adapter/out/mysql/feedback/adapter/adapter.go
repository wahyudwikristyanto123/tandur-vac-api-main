package adapter

import (
	"tandur.com/src/adapter/out/mysql/feedback/repository"
	"tandur.com/src/adapter/out/mysql/feedback/util"
	"tandur.com/src/domain"
)

type FeedbackAdapter struct {
	repo *repository.FeedbackBaseRepository
}

func NewFeedbackAdapter(repo *repository.FeedbackBaseRepository) FeedbackBaseAdapter {
	return &FeedbackAdapter{repo: repo}
}

func (adapter *FeedbackAdapter) GetFeedbackByToken(token string) (*[]domain.Feedback, error) {
	results, err := (*adapter.repo).GetFeedbackByToken(token)
	if err != nil {
		return nil, err
	}
	var data []domain.Feedback = []domain.Feedback{}
	i := 0
	for i < (len(*results)) {
		data = append(data, util.EntityAdapterToDomain((*results)[i]))
		i++
	}
	return &data, nil
}

func (adapter *FeedbackAdapter) GetFeedbackById(id int64) (*domain.Feedback, error) {
	result, err := (*adapter.repo).GetFeedbackById(id)
	if err != nil {
		return nil, err
	}
	data := util.EntityAdapterToDomain((*result))
	return &data, nil
}

func (adapter *FeedbackAdapter) AddFeedback(data domain.Feedback) (int64, error) {
	return (*adapter.repo).InsertFeedback(util.DomainToEntityAdapter(data))
}

func (adapter *FeedbackAdapter) UpdateFeedback(data domain.Feedback) error {
	return (*adapter.repo).UpdateFeedback(util.DomainToEntityAdapter(data))
}

func (adapter *FeedbackAdapter) RemoveFeedback(id int64) error {
	return (*adapter.repo).DeleteFeedback(id)
}
