package controller

import (
	"github.com/go-tron/example/config"
	"github.com/go-tron/example/entity"
	"github.com/go-tron/example/model"
	"github.com/go-tron/example/service"
	"github.com/go-tron/iris/baseContext"
	"github.com/kataras/iris/v12"
	"strconv"
)

var AuthController = authController{}

type authController struct {
}

func (c *authController) Route(p iris.Party) {
	p.Post("/login", baseContext.ReturnHandler(c.Login))
}

func (c *authController) Login(ctx *baseContext.Context) (interface{}, error) {
	var params model.LoginReq
	if err := ctx.JSONReqBody(&params); err != nil {
		return nil, err
	}

	//after validate phone code
	user, err := service.UserService.Login(ctx, &entity.User{Phone: params.Phone})
	if err != nil {
		return nil, err
	}

	token, err := config.UserToken.Create(strconv.FormatInt(user.UserId, 10), 0)
	if err != nil {
		return nil, err
	}

	return &model.LoginRes{
		Token:  token,
		UserId: user.UserId,
		Phone:  user.Phone,
	}, nil
}
