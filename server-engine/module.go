package sample

import (
	"github.com/jhseong7/gimbap"
	"github.com/jhseong7/gimbap-sample/server-engine/food"
	"github.com/jhseong7/gimbap-sample/server-engine/salt"
)

var AppModuleGin = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "AppModuleGin",
	SubModules: []*gimbap.Module{
		salt.SaltModuleGin,
		food.FoodModuleGin,
	},
})

var AppModuleEcho = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "AppModuleEcho",
	SubModules: []*gimbap.Module{
		salt.SaltModuleEcho,
		food.FoodModuleEcho,
	},
})

var AppModuleFiber = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "AppModuleFiber",
	SubModules: []*gimbap.Module{
		salt.SaltModuleFiber,
		food.FoodModuleFiber,
	},
})
