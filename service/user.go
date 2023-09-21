package service

import (
	"context"
	"github.com/go-tron/example/config"
	"github.com/go-tron/example/entity"
	"github.com/go-tron/example/subscriber"
	"github.com/go-tron/local-time"
	"github.com/go-tron/nsq"
)

var UserService = userService{}

type userService struct {
}

func (s *userService) Login(ctx context.Context, user *entity.User) (*entity.User, error) {
	if err := config.DB.FindOne(
		user,
		config.DB.WithSelect("userId,phone,status"),
		config.DB.WithWhere("phone = ?", user.Phone),
	); err != nil {
		if config.DB.IsRecordNotFoundError(err) {
			if err = s.Register(ctx, user); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	if user.Status != 1 {
		return nil, config.ErrorUserStatus
	}
	return user, nil
}

func (s *userService) Register(ctx context.Context, user *entity.User) error {
	user.Status = 1
	user.RegisterAt = localTime.Now()

	if err := config.DB.Create(user); err != nil {
		return err
	}
	config.NsqProducer.SendSync(subscriber.UserRegisterTopic, subscriber.UserRegisterReq{
		UserId:    user.UserId,
		Phone:     user.Phone,
		ChannelId: user.ChannelId,
	}, nsq.WithCtx(ctx))
	return nil
}
