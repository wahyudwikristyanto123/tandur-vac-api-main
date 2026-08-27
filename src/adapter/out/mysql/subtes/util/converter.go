package util

import (
	"strconv"
	"strings"
	"time"

	"tandur.com/src/adapter/out/mysql/subtes/entity"
	"tandur.com/src/domain"
)

func formatLocalTime(val time.Time) time.Time {
	splitted := strings.Split(val.String(), " ")

	year, _ := strconv.Atoi(strings.Split(splitted[0], "-")[0])
	month, _ := strconv.Atoi(strings.Split(splitted[0], "-")[1])
	day, _ := strconv.Atoi(strings.Split(splitted[0], "-")[2])

	hour, _ := strconv.Atoi(strings.Split(splitted[1], ":")[0])
	minute, _ := strconv.Atoi(strings.Split(splitted[1], ":")[1])
	second, _ := strconv.Atoi(strings.Split(splitted[1], ":")[2])

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
}

func EntityAdapterToDomain(data entity.SubtesMySql) domain.Subtes {
	return domain.Subtes{
		ID:                 data.ID,
		AssessmentCenterId: data.AssessmentCenterId,
		Title:              data.Title,
		StartDate:          formatLocalTime(data.StartDate.Time),
		EndDate:            formatLocalTime(data.EndDate.Time),
		Type:               data.Type,
		ToolType:           data.ToolsType,
		Index:              data.OrderSchedule.Int64,
		URLProctoring:      data.URLProctoring.String,
		Status:             data.Status.String,
		Duration:           data.Duration,
	}
}
