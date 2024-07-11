package salt

import (
	"github.com/jhseong7/gimbap"
)

var SaltModuleFiber = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "SaltModule",
	Providers: []gimbap.ProviderDefinition{
		*SaltProvider,
	},
	Controllers: []gimbap.ControllerDefinition{
		*SaltControllerFiberProvider,
	},
})
