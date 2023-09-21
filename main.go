package main

import (
	"github.com/go-tron/example/authorize"
	"github.com/go-tron/example/config"
	"github.com/go-tron/example/controller"
	"github.com/go-tron/example/subscriber"
	"github.com/go-tron/iris/anyError"
	"github.com/go-tron/iris/baseContext"
	"github.com/go-tron/iris/health"
	"github.com/go-tron/iris/recover"
	"github.com/go-tron/iris/requestLogger"
	"github.com/go-tron/iris/response"
	"github.com/go-tron/iris/trace"
	"github.com/go-tron/logger"
	"github.com/go-tron/nsq"
	"github.com/go-tron/tracer"
	"github.com/kataras/iris/v12"
)

func main() {
	_, closer := tracer.NewJaegerWithConfig(config.Config)
	defer closer.Close()

	nsq.NewConsumerWithConfig(config.Config, subscriber.UserRegister)

	rl := requestLogger.New(
		logger.NewZapWithConfig(config.Config, "request", "debug"),
		requestLogger.WithContextKeys(authorize.AuthorizeId),
		requestLogger.WithUserAgent(true),
		requestLogger.WithSessionKeys("accessChannelId"),
	)

	baseContext.New(
		config.Config.GetString("application.env"),
		rl,
		baseContext.WithApplicationName(config.Config.GetString("application.name")),
		baseContext.WithInternal(config.Config.GetBool("application.internal")),
		baseContext.WithResponse(response.New()),
		baseContext.WithViewError("error.html"),
	)

	app := iris.New()
	app.HandleDir("/", "./public")
	app.RegisterView(iris.HTML("./public", ".html"))

	health.New(app, health.WithHandlers(config.BasicAuth))
	app.Use(recover.New())
	app.Use(trace.New())
	app.Use(config.Sessions.Handler())
	app.Use(rl.Handler())
	app.OnAnyErrorCode(anyError.New())

	app.PartyFunc("/", controller.IndexController.Route)

	app.Use(authorize.Authorize.Handler())
	app.PartyFunc("/api/auth", controller.AuthController.Route)
	app.PartyFunc("/api/address", controller.AddressController.Route)

	app.Run(iris.Addr(":"+config.Config.GetString("application.port")),
		iris.WithoutBodyConsumptionOnUnmarshal,
		iris.WithoutPathCorrection,
	)
}
