/*
*
This is an example app to demonstrate how to initialize a microservice with getting dependencies from the DI container.
*/
package main

import (
	"github.com/jhseong7/gimbap"
	ms "github.com/jhseong7/gimbap-sample/microservice-sample"
	ss "github.com/jhseong7/gimbap-sample/server-engine"
	fiber_engine "github.com/jhseong7/gimbap/engine/fiber"
)

func main() {
	app := gimbap.CreateApp(gimbap.AppOption{
		AppName:      "Test",
		AppModule:    ss.AppModuleFiber,
		ServerEngine: fiber_engine.NewFiberHttpEngine(),
	})

	// Add the microservice to the app
	app.AddMicroServices(ms.MicroserviceProvider, ms.OtherSampleMicroServiceProvider)

	app.Run()
}
