package subscriber

import (
	"context"
	"encoding/json"
	baseError "github.com/go-tron/base-error"
	"github.com/go-tron/nsq"
)

const UserRegisterTopic = "user-register"

type UserRegisterReq struct {
	UserId    int64  `json:"userId"`
	Phone     string `json:"string"`
	ChannelId int    `json:"channelId"`
}

var UserRegister = &nsq.ConsumerConfig{
	Topic:           UserRegisterTopic,
	BackoffDisabled: true,
	Handler: func(ctx context.Context, msg string, finished bool) error {
		m := &UserRegisterReq{}
		if err := json.Unmarshal([]byte(msg), m); err != nil {
			return err
		}
		return baseError.System("123", "4456")
	},
}
