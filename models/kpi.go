package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

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
	Thailand string `json:"thailand"`
}

type UserKPI struct {
	Cdate string `json:"cdate"`
	Mcu   string `json:"mcu"`
	Avg   string `json:"avg"`
	Uu    string `json:"uu"`
	Nru   string `json:"nru"`
}

type SaleKPI struct {
	Cdate    string  `json:"cdate"`
	Rev      string  `json:"rev"`
	Arppu    string  `json:"arppu"`
	Bu       string  `json:"bu"`
	Prate    float32 `json:"prate"`
	Charge   float32 `json:"charge"`
	Chargeuu float32 `json:"chargeuu"`
}

type SaleItemKPI struct {
	Cdate    string  `json:"cdate"`
	Itemid   string  `json:"itemid"`
	Itemname string  `json:"itemname"`
	Count    float32 `json:"count"`
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
			" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
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
func (k *UserKPI) GetUserKPI(from, to, country, kind, radio, kindCalendar string) ([]UserKPI, error) {
	var listKpi []UserKPI

	var sql string
	var err error
	o := orm.NewOrm()

	fmt.Println("input: ", from, to, country, kind, radio, kindCalendar)

	// make query
	sCounty := ""
	if country != "all" {
		sCounty = " and territory = ? "
	}

	// day
	if kindCalendar == "day" {
		sql = " select date cdate " +
			" , uu_d uu" +
			" , nru_d nru" +
			" , mcu_d mcu" +
			" , avg_d avg" +
			" from kpi " +
			" where date >= ? and date <=  ? " +
			sCounty +
			" order by 1;"

		if country != "all" {
			_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
		} else {
			_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
		}

	} else {
		// others
		setD := "yyyy-mm-dd"
		if kindCalendar == "month" {
			setD = "yyyy-mm"
		}
		sql = " select to_char(date_trunc('" + kindCalendar + "',date), '" + setD + "' ) cdate" +
			" , sum(uu_d) uu" +
			" , sum(nru_d) nru" +
			" , sum(mcu_d) mcu" +
			" , sum(avg_d) avg" +
			" from kpi " +
			" where date >= ? and date <=  ? " +
			sCounty +
			//" group by date_trunc('?',date)" +
			" group by cdate " +
			" order by 1;"

		if country != "all" {
			_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
		} else {
			_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
		}

	}

	return listKpi, err
}

// GetSaleKPI ...
func (k *SaleKPI) GetSaleKPI(from, to, country, kind, radio, kindCalendar string) ([]SaleKPI, error) {
	var listKpi []SaleKPI

	var sql string
	var err error
	o := orm.NewOrm()

	fmt.Println("input: ", from, to, country, kind, radio, kindCalendar)

	// make query
	sCounty := ""
	if country != "all" {
		sCounty = " and territory = ? "
	}

	setRev := " , rev_t rev"
	if country == "KOREA" {
		setRev = " , rev_d rev"
	}

	// day
	if kindCalendar == "day" {
		sql = " select date cdate " +
			//" , rev_d rev" +
			//" , rev_t rev" +
			setRev +
			" , arppu_d arppu" +
			" , bu " +
			" , pur_d prate" +
			// " , rev_t rev_t" +
			// " , rev_rate rev_rate" +
			" , coalesce(charge_d, 0) charge " +
			" , coalesce(charge_unique_user_d, 0) chargeuu " +
			" from kpi " +
			" where date >= ? and date <=  ? " +
			sCounty +
			" order by 1;"

		if country != "all" {
			_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
		} else {
			_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
		}

	} else {
		// others
		setD := "yyyy-mm-dd"
		if kindCalendar == "month" {
			setD = "yyyy-mm"
		}

		setP := " , sum(pur_d) / 7 prate"
		if kindCalendar == "month" {
			setP = " , sum(pur_d) / 30 prate"
		}

		setRev = " , sum(rev_t) rev"
		if country == "KOREA" {
			setRev = " , sum(rev_d) rev"
		}

		sql = " select to_char(date_trunc('" + kindCalendar + "',date), '" + setD + "' ) cdate" +
			//" , sum(rev_d) rev" +
			//" , sum(rev_t) rev" +
			setRev +
			" , sum(arppu_d) arppu" +
			" , sum(bu) bu" +
			setP +
			" , coalesce(sum(charge_d), 0) charge " +
			" , coalesce(sum(charge_unique_user_d), 0) chargeuu " +
			" from kpi " +
			" where date >= ? and date <=  ? " +
			sCounty +
			//" group by date_trunc('?',date)" +
			" group by cdate " +
			" order by 1;"

		if country != "all" {
			_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
		} else {
			_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
		}

	}

	return listKpi, err
}

// GetSaleItemKPI ...
func (k *SaleItemKPI) GetSaleItemKPI(from, to, country, kind, radio, kindCalendar string) ([]SaleItemKPI, error) {
	var listKpi []SaleItemKPI

	var sql string
	var err error
	o := orm.NewOrm()

	fmt.Println("input: ", from, to, country, kind, radio, kindCalendar)

	// make query
	sCounty := ""
	if country != "all" {
		sCounty = " and country = ? "
	}

	// day
	if kindCalendar == "day" {
		sql = " select  " +
			"  \"ExternalItemID\" itemid " +
			" , external_item_name_eng itemname " +
			" , sum(sales_count) count" +
			" from item_kpi " +
			" where date >= ? and date <=  ? " +
			sCounty +
			" group by itemid, itemname " +
			" order by 3 desc " +
			" limit 50 "

		if country != "all" {
			_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
		} else {
			_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
		}

	} else {
		// others
		/*
			setD := "yyyy-mm-dd"
			if kindCalendar == "month" {
				setD = "yyyy-mm"
			}
		*/

		sql = //" select to_char(date_trunc('" + kindCalendar + "',date), '" + setD + "' ) cdate" +
			" select " +
				"  \"ExternalItemID\" itemid " +
				" , external_item_name_eng itemname " +
				" , sum(sales_count) count" +
				" from item_kpi " +
				" where date >= ? and date <=  ? " +
				sCounty +
				//" group by date_trunc('?',date)" +
				" group by itemid, itemname " +
				" order by 3 desc " +
				" limit 50 "

		if country != "all" {
			_, err = o.Raw(sql, from, to, country).QueryRows(&listKpi)
		} else {
			_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
		}

	}

	return listKpi, err
}
