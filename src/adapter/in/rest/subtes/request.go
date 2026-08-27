package subtes

type UpdateRequest struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type SubmitRequest struct {
	Id     int64  `json:"id"`
	Result string `json:"result"`
}
