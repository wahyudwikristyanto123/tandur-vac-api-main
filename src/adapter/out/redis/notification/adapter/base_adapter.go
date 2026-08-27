package adapter

type RedisNotificationBaseAdapter interface {
	SetMessage(token string, message string) error
	GetByToken(token string) (*string, error)
}
