package models

import (
	"log"
	"os"

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
	)

	orm.RegisterDriver("postgres", orm.DRPostgres)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		//beego.Error("Error loading .env file")
	}

	DBHOST := os.Getenv("DBHOST")

	orm.RegisterDataBase("default", "postgres", DBHOST)
}
