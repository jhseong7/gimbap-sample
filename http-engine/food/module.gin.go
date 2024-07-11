package food

import "github.com/jhseong7/gimbap"

var FoodModuleGin = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "FoodModuleGin",
	Providers: []gimbap.ProviderDefinition{
		*FoodProvider,
	},
	Controllers: []gimbap.ControllerDefinition{
		*FoodControllerGinProvider,
	},
})
