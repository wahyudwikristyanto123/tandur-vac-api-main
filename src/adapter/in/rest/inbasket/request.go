package inbasket

type ReplyRequest struct {
	ParentId int64  `json:"parent_id"`
	Message  string `json:"message"`
}

type FullReplyRequest struct {
	ParentId    int64    `json:"parent_id"`
	Ccs         []string `json:"ccs"`
	Attachments []string `json:"attachments"`
	Message     string   `json:"message"`
}

type UpdateRequest struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type UpdateDraftRequest struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type UpdateFullDraftRequest struct {
	Id          int64    `json:"id"`
	Message     string   `json:"message"`
	Ccs         []string `json:"ccs"`
	Attachments []string `json:"attachments"`
}
