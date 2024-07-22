package main

import (
	"github.com/jhseong7/gimbap"
	sample "github.com/jhseong7/gimbap-sample/server-engine"
	"github.com/jhseong7/gimbap/engine"
	"github.com/labstack/echo/v4"
)

func main() {
	app := gimbap.CreateApp(gimbap.AppOption{
		AppName:      "SampleAppEcho",
		AppModule:    sample.AppModuleEcho,
		ServerEngine: engine.NewEchoHttpEngine(),
	})

	app.AddMiddleware(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			return next(ctx)
		}
	})

	app.Run()
}
