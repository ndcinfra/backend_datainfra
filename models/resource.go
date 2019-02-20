package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Resource ...
type Resource struct {
	ID              int64  `orm:"column(ID);auto;pk" json:"id"` // id
	Sheet           string `orm:"size(1000);null" json:"sheet"`
	Memo            string `orm:"size(2000);null" json:"memo"`
	Seha            string `orm:"size(2000);null" json:"seha"`
	Sylvi           string `orm:"size(2000);null" json:"sylvi"`
	Yuri            string `orm:"size(2000);null" json:"yuri"`
	Misteltein      string `orm:"size(2000);null" json:"misteltein"`
	Jay             string `orm:"size(2000);null" json:"jay"`
	Harpy           string `orm:"size(2000);null" json:"harpy"`
	Levia           string `orm:"size(2000);null" json:"levia"`
	Nata            string `orm:"size(2000);null" json:"nata"`
	Tina            string `orm:"size(2000);null" json:"tina"`
	Violet          string `orm:"size(2000);null" json:"violet"`
	Wolfgang        string `orm:"size(2000);null" json:"wolfgang"`
	Soma            string `orm:"size(2000);null" json:"soma"`
	Luna            string `orm:"size(2000);null" json:"luna"`
	MaleAccessory   string `orm:"column(MaleAccessory);size(2000);null" json:"maleaccessory"`
	FemaleAccessory string `orm:"column(FemaleAccessory);size(2000);null" json:"femaleaccessory"`
}

// AddResource ...
func AddResource(r Resource) (int64, error) {
	result, err := orm.NewOrm().Insert(&r)
	if err != nil {
		beego.Error("insert into resource: ", err)
		return -1, err
	}

	return result, nil
}

// GetResourceAll ...
func GetResources() ([]Resource, error) {
	var resource []Resource

	o := orm.NewOrm()
	sql := "SELECT " +
		" \"ID\" , " +
		" Sheet, " +
		" Memo, " +
		" Seha, " +
		" Sylvi, " +
		" Yuri, " +
		" Misteltein, " +
		" Jay, " +
		" Harpy, " +
		" Levia, " +
		" Nata, " +
		" Tina, " +
		" Violet, " +
		" Wolfgang, " +
		" Soma, " +
		" Luna, " +
		" \"MaleAccessory\", " +
		" \"FemaleAccessory\" " +
		" FROM \"resource\"  " +
		" ORDER BY \"ID\" ASC"

	_, err := o.Raw(sql).QueryRows(&resource)
	if err != nil {
		return resource, err
	}

	return resource, nil
}

// GetResourceDetail ...
func GetResourceDetail(id int) (Resource, error) {
	var resource Resource

	o := orm.NewOrm()
	sql := "SELECT " +
		" \"ID\" , " +
		" Sheet, " +
		" Memo, " +
		" Seha, " +
		" Sylvi, " +
		" Yuri, " +
		" Misteltein, " +
		" Jay, " +
		" Harpy, " +
		" Levia, " +
		" Nata, " +
		" Tina, " +
		" Violet, " +
		" Wolfgang, " +
		" Soma, " +
		" Luna, " +
		" \"MaleAccessory\", " +
		" \"FemaleAccessory\" " +
		" FROM \"resource\"  " +
		" WHERE \"ID\" = ?"

	err := o.Raw(sql, id).QueryRow(&resource)
	if err != nil {
		return resource, err
	}

	return resource, nil
}

// UpdateResource ...
func UpdateResource(r Resource) (Resource, error) {
	o := orm.NewOrm()

	if _, err := o.Update(&r); err != nil {
		return Resource{}, err
	}

	return r, nil
}

// DeleteResource ...
func DeleteResource(id int64) error {

	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM \"resource\" WHERE \"ID\" = ?", id).Exec()

	if err != nil {
		return err
	}

	return nil

}
