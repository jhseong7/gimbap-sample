package food

import "github.com/jhseong7/gimbap"

var FoodModuleFiber = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "FoodModuleFiber",
	Providers: []gimbap.ProviderDefinition{
		*FoodProvider,
	},
	Controllers: []gimbap.ControllerDefinition{
		*FoodControllerFiberProvider,
	},
})
