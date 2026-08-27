package adapter

import (
	"tandur.com/src/util"
)

type RedisNotificationAdapter struct{}

func NewRedisNotificationAdapter() RedisNotificationBaseAdapter {
	return &RedisNotificationAdapter{}
}

func (adapter *RedisNotificationAdapter) SetMessage(token string, message string) error {
	err := util.GetRedis().Set(util.GetRedisContext(), token, message, util.GetDefaultRedisExpiration()).Err()
	if err != nil {
		return err
	}
	return nil
}

func (adapter *RedisNotificationAdapter) GetByToken(token string) (*string, error) {
	val, err := util.GetRedis().Get(util.GetRedisContext(), token).Result()
	if err != nil {
		return nil, err
	}
	return &val, nil
}
