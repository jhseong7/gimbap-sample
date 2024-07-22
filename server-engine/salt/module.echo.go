package salt

import (
	"github.com/jhseong7/gimbap"
)

var SaltModuleEcho = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "SaltModule",
	Providers: []*gimbap.Provider{
		SaltProvider,
	},
	Controllers: []*gimbap.Controller{
		SaltControllerEchoProvider,
	},
})
