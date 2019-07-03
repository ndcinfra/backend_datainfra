package models

import "github.com/astaxie/beego/orm"

// Kpi ...
// for table
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

// KpiGraph ...
type KpiGraph struct {
	Cdate    string `json:"cdate"`
	China    string `json:"china"`
	Japan    string `json:"japan"`
	Korea    string `json:"korea"`
	Namerica string `json:"namerica"`
	Taiwan   string `json:"taiwan"`
	Total    string `json:"total"`
}

// GetKPI ...
func (k *Kpi) GetKPI(from, to, country, kind, radio string) ([]Kpi, []KpiGraph, error) {
	var listKpi []Kpi
	var gListKpi []KpiGraph

	var err error
	var sql string

	o := orm.NewOrm()

	var scol string
	switch radio {
	case "rev":
		scol = "rev_d"
		break
	case "avg":
		scol = "avg_d"
		break
	case "mcu":
		scol = "mcu_d"
		break
	case "uu":
		scol = "uu_d"
		break
	case "nru":
		scol = "nru_d"
		break
	default:
		scol = "rev_d"
	}

	if kind == "graph" {
		sql = " select cdate " +
			" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
			" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
			" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
			" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
			" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
			" ,sum(rev) TOTAL " +
			" from ( " +
			"			select " +
			"				date as cdate " +
			"				, territory  " +
			"				, " + scol + " as rev  " +
			" 		from kpi where date >= ? and date <=  ? " +
			"			order by date, territory asc " +
			"		) a " +
			" group by cdate" +
			" order by 1;"
		_, err = o.Raw(sql, from, to).QueryRows(&gListKpi)
	}

	if kind == "table" {
		sql = "SELECT territory, to_char(date,'YYYY-MM-DD') as date, mcu_d as mcu, " +
			"avg_d as avg, uu_d as uu, nru_d as nru, " +
			"rev_d as rev, rev_ur_d as rev_ur, rev_pc_d as rev_pc, pu_d as pu, pur_d as pur, " +
			"arppu_d as arppu, dt, mrppu_d as mrppu, rev_t, rev_rate, bu " +
			"FROM kpi WHERE date >= ? and date <= ? "
		_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
	}

	return listKpi, gListKpi, err
}

// GetUserKPI ...
func (k *Kpi) GetUserKPI(from, to, country, kind, radio, kindCalendar string) ([]Kpi, error) {
	var sql string

	// day
	if kindCalendar == 'day' {

	}else{
		
	}
	// others
}
