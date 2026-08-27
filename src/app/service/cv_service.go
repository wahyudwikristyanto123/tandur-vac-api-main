package service

import (
	"tandur.com/src/adapter/out/mysql/cv/adapter"
	usecase "tandur.com/src/app/use_case"
	"tandur.com/src/domain"
)

type CvService struct {
	adapter *adapter.CvBaseAdapter
}

func NewCvService(adapter *adapter.CvBaseAdapter) usecase.CvUseCase {
	return &CvService{
		adapter: adapter,
	}
}

func (service *CvService) GetByToken(token string) (*domain.Cv, error) {
	data, err := (*service.adapter).GetCvByToken(token)
	if err != nil {
		return nil, err
	}
	if len(*data) < 1 {
		return nil, nil
	}
	return &(*data)[0], nil
}

func (service *CvService) SubmitCv(data domain.Cv) error {
	results, err := (*service.adapter).GetCvByToken(data.Token)
	if err != nil {
		return err
	}
	finalData := domain.Cv{
		Token:              data.Token,
		CmdId:              data.CmdId,
		AssessmentCenterId: data.AssessmentCenterId,
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
	if len(*results) < 1 {
		_, err = (*service.adapter).AddCv(finalData)
		if err != nil {
			return err
		}
	} else {
		finalData.Id = (*results)[0].Id
		err = (*service.adapter).UpdateCv(finalData)
		if err != nil {
			return err
		}
	}

	return nil
}
