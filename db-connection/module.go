package dbconnection

import (
	"github.com/jhseong7/gimbap"
	"github.com/jhseong7/gimbap-sample/db-connection/db"
)

var DbConnectionModule = gimbap.DefineModule(gimbap.ModuleOption{
	Name: "DbConnection",
	Providers: []gimbap.ProviderDefinition{
		*db.GormProvider,
	},
})
