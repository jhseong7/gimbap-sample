package salt

import (
	"github.com/jhseong7/gimbap/controller"
	"github.com/jhseong7/gimbap/core"
	"github.com/jhseong7/gimbap/provider"
)

var SaltModuleEcho = core.DefineModule(core.ModuleOption{
	Name: "SaltModule",
	Providers: []provider.ProviderDefinition{
		*SaltProvider,
	},
	Controllers: []controller.ControllerDefinition{
		*SaltControllerEchoProvider,
	},
})