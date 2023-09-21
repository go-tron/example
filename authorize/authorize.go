package authorize

import (
	"github.com/go-tron/example/config"
	"github.com/go-tron/iris/authorize/bearerToken"
	"github.com/go-tron/iris/baseContext"
	"regexp"
	"strconv"
)

const (
	AuthorizeId   = "userId"
	AuthorizeUser = "user"
)

type User struct {
	UserId   int64  `gorm:"primary_key" json:"userId,string"`
	DoryUid  string `json:"doryUid"`
	Username string `json:"username"`
}

func GetUserId(ctx *baseContext.Context) int64 {
	return ctx.Values().GetInt64Default(AuthorizeId, 0)
}
func GetUser(ctx *baseContext.Context) *User {
	if v := ctx.Values().Get("user"); v != nil {
		user, ok := v.(*User)
		if ok {
			return user
		}
	}
	return nil
}

type authorize struct {
}

func (s authorize) Handler(ctx *baseContext.Context, token string) error {
	userId := ctx.GetSession().GetInt64Default(AuthorizeId, 0)
	if userId == 0 {
		id, _ := config.UserToken.Verify(token)
		if id != "" {
			userId, _ = strconv.ParseInt(id, 10, 64)
		}
	}
	if userId == 0 {
		return config.ErrorAuthorize
	}

	var user = User{
		UserId: userId,
	}
	if err := config.DB.FindById(
		&user,
		config.DB.WithTable("user"),
		config.DB.WithSelect("userId,name"),
		config.DB.WithErrorNotFound(config.ErrorUserNotExists),
	); err != nil {
		return err
	}

	ctx.Values().Set(AuthorizeId, userId)
	ctx.Values().Set(AuthorizeUser, &user)
	return nil
}

var Authorize = bearerToken.New(
	&authorize{},
	bearerToken.WithIgnorePaths(
		regexp.MustCompile("^/api/auth"),
	),
)
