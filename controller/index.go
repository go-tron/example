package controller

import (
	baseError "github.com/go-tron/base-error"
	"github.com/go-tron/example/config"
	"github.com/go-tron/iris/baseContext"
	localTime "github.com/go-tron/local-time"
	"github.com/go-tron/logger"
	"github.com/kataras/iris/v12"
)

var IndexController = indexController{}

type indexController struct {
}

func (c *indexController) Route(p iris.Party) {
	p.Get("/", baseContext.Handler(c.Index))
	p.Get("/error", baseContext.Handler(c.Error))
	p.Get("/returnError", baseContext.ReturnHandler(c.ReturnError))
	p.Get("/panic", baseContext.ReturnHandler(c.Panic))
}

func (c *indexController) Index(ctx *baseContext.Context) {
	config.OrderLogger.Error("hi", logger.NewField("aaa", "bbb"))
	session := ctx.GetSession()
	session.Set("lastVisit", localTime.Now())
	session.Increment("visits", 1)
	ctx.ViewData("Visits", session.GetIntDefault("visits", 1))
	ctx.View("index.html")
}

func (c *indexController) Error(ctx *baseContext.Context) {
	var err = config.ErrorSystem("hahaha")
	ctx.ViewData("Message", err.Error())
	ctx.View("error.html")
}

func (c *indexController) ReturnError(ctx *baseContext.Context) (interface{}, error) {
	ctx.AddLogField("test", "hello")
	return "hi", baseError.SystemStack("9998", "系统错误222", 3)
}

func (c *indexController) Panic(ctx *baseContext.Context) (interface{}, error) {
	var a = make([]int, 0)
	a[1] = 1
	return "hi", nil
}
