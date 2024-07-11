package sample

import (
	"github.com/jhseong7/gimbap"
	"github.com/jhseong7/gimbap-sample/http-engine/food"
	"github.com/jhseong7/gimbap-sample/http-engine/salt"
	"github.com/jhseong7/gimbap/module"
)

var AppModuleGin = module.DefineModule(gimbap.ModuleOption{
	Name: "AppModuleGin",
	SubModules: []gimbap.Module{
		*salt.SaltModuleGin,
		*food.FoodModuleGin,
	},
})

var AppModuleEcho = module.DefineModule(gimbap.ModuleOption{
	Name: "AppModuleEcho",
	SubModules: []module.Module{
		*salt.SaltModuleEcho,
		*food.FoodModuleEcho,
	},
})

var AppModuleFiber = module.DefineModule(gimbap.ModuleOption{
	Name: "AppModuleFiber",
	SubModules: []module.Module{
		*salt.SaltModuleFiber,
		*food.FoodModuleFiber,
	},
})
