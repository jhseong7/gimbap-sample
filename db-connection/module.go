package dbconnection

import (
	"github.com/jhseong7/gimbap-sample/db-connection/db"
	"github.com/jhseong7/gimbap/core"
	"github.com/jhseong7/gimbap/provider"
)

var DbConnectionModule = core.DefineModule(core.ModuleOption{
	Name: "DbConnection",
	Providers: []provider.ProviderDefinition{
		*db.GormProvider,
	},
})
