package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Indonesia ...
type Indonesia struct {
	ID       int       `orm:"column(ID);pk;auto" json:"id"`                 // max 100 letters
	Email    string    `orm:"size(100);unique" json:"email"`                // max 100 letters
	IP       string    `orm:"column(IP);size(20);null" json:"ip"`           //
	CreateAt time.Time `orm:"type(datetime);auto_now_add" json:"create_at"` // first save
}

// AddIndonesiaData ...
func AddIndonesiaData(i Indonesia) (int64, error) {
	result, err := orm.NewOrm().Insert(&i)
	if err != nil {
		//beego.Error("insert into indonesia: ", err)
		return -1, err
	}

	return result, nil
}
