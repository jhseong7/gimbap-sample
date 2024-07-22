package db

import (
	logger "github.com/jhseong7/ecl"
	"github.com/jhseong7/gimbap"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	DbClientOption struct {
		DatabaseType string
		DSN          string
	}
)

func getMySqlClient(option DbClientOption) *gorm.DB {
	db, err := gorm.Open(mysql.Open(option.DSN), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func getPostgresClient(option DbClientOption) *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: option.DSN,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db

}

func getSqlLiteClient(option DbClientOption) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(option.DSN), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db

}

func NewGormClient(option DbClientOption) (db *gorm.DB) {
	var log = logger.NewLogger(logger.LoggerOption{
		Name: "GOrmClient",
	})

	switch option.DatabaseType {
	case "postgres":
		db = getPostgresClient(option)
	case "sqlite":
		db = getSqlLiteClient(option)
	case "mysql":
		db = getMySqlClient(option)
	default:
		log.Panicf("Database type %s is not supported", option.DatabaseType)
	}
	return
}

var GormProvider = gimbap.DefineProvider(gimbap.ProviderOption{
	Name:         "GormProvider",
	Instantiator: NewGormClient,
})
