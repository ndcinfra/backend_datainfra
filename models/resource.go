package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Resource ...
type Resource struct {
	ID        int64  `orm:"column(ID);auto;pk" json:"id"` // id
	Character string `orm:"size(1000);null" json:"character"`
	ImgURL    string `orm:"column(ImgURL);size(1000);null" json:"imgurl"`
}

func AddResource(r Resource) (int64, error) {
	result, err := orm.NewOrm().Insert(&r)
	if err != nil {
		beego.Error("insert into resource: ", err)
		return -1, err
	}
	fmt.Println("ID: ", result)
	/*
		o := orm.NewOrm()

		sql := "INSERT INTO \"resource\" " +
			"(\"ImgURL\") " +
			"VALUES ($1)"

		result, err := o.Raw(sql, r.ImgURL).Exec()
		if err != nil {
			beego.Error("insert into resource: ", err)
			return -1, err
		}

		ID, _ := result..LastInsertId()
		fmt.Println("ID: ", ID)
	*/
	return result, nil
}

// GetResourceAll ...
func GetResourceAll() ([]Resource, error) {
	var resource []Resource

	o := orm.NewOrm()
	sql := "SELECT " +
		" \"ID\" , " +
		" Character, " +
		" \"ImgURL\" " +
		" FROM \"resource\"  "

	_, err := o.Raw(sql).QueryRows(&resource)
	if err != nil {
		return resource, err
	}

	return resource, nil
}
