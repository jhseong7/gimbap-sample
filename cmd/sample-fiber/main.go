package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jhseong7/gimbap"
	sample "github.com/jhseong7/gimbap-sample/server-engine"
	"github.com/jhseong7/gimbap/engine"
)

func main() {
	app := gimbap.CreateApp(gimbap.AppOption{
		AppName:   "SampleAppFiber",
		AppModule: sample.AppModuleFiber,
		ServerEngine: engine.NewFiberHttpEngine(engine.FiberHttpEngineOption{
			FiberConfig: fiber.Config{
				AppName: "SampleAppFiber",
			},
		}),
	})

	// Example of global middleware
	app.AddMiddleware(func(ctx *fiber.Ctx) (e error) {
		ctx.Set("Access-Control-Allow-Origin", "*")
		ctx.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		return
	})

	app.Run()
}
