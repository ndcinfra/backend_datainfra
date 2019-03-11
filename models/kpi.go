package models

import "github.com/astaxie/beego/orm"

// Kpi ...
type Kpi struct {
	Territory string `json:"territory"` // territory
	Date      string `json:"date"`
	Mcu       string `json:"mcu"`
	Avg       string `json:"avg"`
	Uu        string `json:"uu"`
	Nru       string `json:"nru"`
	Rev       string `json:"rev"`
	RevUr     string `json:"rev_ur"`
	RevPc     string `json:"rev_pc"`
	Pu        string `json:"pu"`
	Pur       string `json:"pur"`
	Arppu     string `json:"arppu"`
	Dt        string `json:"dt"`
	Mrppu     string `json:"mrppu"`
	RevT      string `json:"rev_t"`
	RevRate   string `json:"rev_rate"`
	Bu        string `json:"bu"`
}

// GetKPI ...
func (k *Kpi) GetKPI(from, to, country string) ([]Kpi, error) {
	var listKpi []Kpi

	o := orm.NewOrm()
	sql := "SELECT territory, date, mcu_d as mcu, avg_d as avg, uu_d as uu, nru_d as nru, " +
		"rev_d as rev, rev_ur_d as rev_ur, rev_pc_d as rev_pc, pu_d as pu, pur_d as pur, " +
		"arppu_d as arppu, dt, mrppu_d as mrppu, rev_t, rev_rate, bu " +
		"FROM kpi WHERE date >= ? and date <= ? "

	var err error
	if country == "all" {
		sql = sql + " ORDER BY date desc "
		_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
	} else {
		sql = sql + " and territory = ? ORDER BY date desc "
		_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
	}

	return listKpi, err
}
