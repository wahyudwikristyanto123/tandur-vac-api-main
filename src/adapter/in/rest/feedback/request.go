package feedback

type FeedbackRequest struct {
	Token    string `json:"token"`
	AcId     int64  `json:"ac_id"`
	Feedback string `json:"feedback"`
}
