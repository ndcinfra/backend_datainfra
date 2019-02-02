package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql" ...
)

// RegisterDB ...
func RegisterDB() {
	// register model
	orm.RegisterModel(
		new(User),
		new(Resource),
	)

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://gauvkrkt:Et0dCgKDaY2RCu-GW9Az5JtbzAGOLcfv@stampy.db.elephantsql.com:5432/gauvkrkt")
}
