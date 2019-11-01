package models

import (
	"os"

	"github.com/astaxie/beego/logs"

	"github.com/astaxie/beego/orm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// RegisterDB ...
func RegisterDB() {
	// register model
	orm.RegisterModel(
		new(User),
		new(Resource),
		new(Indonesia),
		new(CountryShare),
	)

	orm.RegisterDriver("postgres", orm.DRPostgres)

	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
	}

	DBHOST := os.Getenv("DBHOST")

	orm.RegisterDataBase("default", "postgres", DBHOST)
}
