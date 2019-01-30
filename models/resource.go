package models

// Resource ...
type Resource struct {
	ID     int64  `orm:"column(ID);auto;pk" json:"id"` // id
	ImgURL string `orm:"column(ImgURL);size(1000);null" json:"imgurl"`
}
