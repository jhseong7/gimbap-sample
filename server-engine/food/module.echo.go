package food

import "github.com/jhseong7/gimbap"

var FoodModuleEcho = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "FoodModuleEcho",
	Providers: []*gimbap.Provider{
		FoodProvider,
	},
	Controllers: []*gimbap.Controller{
		FoodControllerEchoProvider,
	},
})
