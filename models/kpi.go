package models

import (
	"fmt"

	"time"

	"github.com/astaxie/beego/orm"
)

//table country_share //NADDIC에서 받는 매출%
type InputDate struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Country      string `json:"country"`
	Kind         string `json:"kind"`         // graph, table
	Radio        string `json:"radio"`        // uu, mcu ...
	KindCalendar string `json:"kindCalendar"` // day, week, month
	Period       string `json:"period"`
}

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
	Vietnam  string `json:"vietnam"`
}

type UserKPI struct {
	Cdate string `json:"cdate"`
	Mcu   string `json:"mcu"`
	Avg   string `json:"avg"`
	Uu    string `json:"uu"`
	Nru   string `json:"nru"`
	Gnru  string `json:"gnru"`
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

type CountryShare struct {
	ID          int64     `orm:"column(ID);auto;pk" json:"id"`                  // id, key
	Territory   string    `orm:"size(1000);null" json:"territory"`              // country
	StartDate   time.Time `orm:"type(datetime);auto_now_add" json:"start_date"` // start date
	EndDate     time.Time `orm:"type(datetime);auto_now_add" json:"end_date"`   // end date
	CountryRate string    `orm:"size(10);null" json:"c_reate"`                  // rate
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

// GetNewKPI ...
func (k *Kpi) GetNewKPI(from, to, country, kind, radio string, period string) ([]Kpi, []KpiGraph, error) {
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
	case "nad":
		scol = "rev_d"
		break
	default:
		scol = "rev_d"
	}

	//나딕 매출 통계 프로세스 시작==================================================================================
	if radio == "nad" {
		//PERIOD SEARCH (1:DAILY, 2:WEEKLY, 3:MONTHLY)
		if period == "2" {

			//WEEKLY SEARCH
			if kind == "graph" {
				sql = " select cdate " +
					" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
					" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
					" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
					" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
					" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
					" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
					" ,sum(case when territory = 'VIETNAM' then rev else 0 end) VIETNAM" +
					" ,sum(rev) TOTAL " +
					" from ( " +
					"           SELECT " +
					"               date(date_trunc('week', kpi.date))+6 as cdate, " +
					"               kpi.territory,  " +
					"               (kpi." + scol + " * (CAST(shar.country_rate AS FLOAT)/100)) AS rev  " +
					"               FROM kpi AS kpi  " +
					"               LEFT OUTER JOIN country_share AS shar " +
					"               ON kpi.territory = shar.territory  " +
					"               AND to_char(kpi.date,'YYYYMMDD') >= to_char(shar.start_date,'YYYYMMDD')  " +
					"               AND to_char(kpi.date,'YYYYMMDD') <= to_char(shar.end_date,'YYYYMMDD') " +
					"               where date(date_trunc('week', kpi.date))+6 >= ? and kpi.date <=  ? " +
					"               ORDER BY kpi.date, kpi.territory asc " +
					"		) a " +
					" group by ROLLUP(cdate)" +
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
		} else if period == "3" {
			if kind == "graph" {

				//MONTHLY SEARCH	//일자로 받은 데이터를 월까지만...
				from_substring := string(from[0:6])
				to_substring := string(to[0:6])

				sql = " select cdate " +
					" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
					" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
					" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
					" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
					" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
					" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
					" ,sum(case when territory = 'VIETNAM' then rev else 0 end) VIETNAM" +
					" ,sum(rev) TOTAL " +
					" from ( " +
					"           SELECT " +
					"               to_char(kpi.date,'YYYY-MM') as cdate, " +
					"               kpi.territory,  " +
					"               (kpi." + scol + " * (CAST(shar.country_rate AS FLOAT)/100)) AS rev  " +
					"               FROM kpi AS kpi  " +
					"               LEFT OUTER JOIN country_share AS shar " +
					"               ON kpi.territory = shar.territory  " +
					"               AND to_char(kpi.date,'YYYYMMDD') >= to_char(shar.start_date,'YYYYMMDD')  " +
					"               AND to_char(kpi.date,'YYYYMMDD') <= to_char(shar.end_date,'YYYYMMDD') " +
					"               where to_char(kpi.date,'YYYYMM') >= ? and to_char(kpi.date,'YYYYMM') <=  ? " +
					"               ORDER BY kpi.date, kpi.territory asc " +
					"		) a " +
					" group by ROLLUP(cdate)" +
					" order by 1;"
				_, err = o.Raw(sql, from_substring, to_substring).QueryRows(&gListKpi)
			}

			if kind == "table" {
				sql = "SELECT territory, to_char(date,'YYYY-MM') as date, mcu_d as mcu, " +
					"avg_d as avg, uu_d as uu, nru_d as nru, " +
					"rev_d as rev, rev_ur_d as rev_ur, rev_pc_d as rev_pc, pu_d as pu, pur_d as pur, " +
					"arppu_d as arppu, dt, mrppu_d as mrppu, rev_t, rev_rate, bu " +
					"FROM kpi WHERE date >= ? and date <= ? "
				_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
			}
		} else {
			if kind == "graph" {

				//DAILY SEARCH
				sql = " select cdate " +
					" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
					" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
					" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
					" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
					" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
					" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
					" ,sum(case when territory = 'VIETNAM' then rev else 0 end) VIETNAM" +
					" ,sum(rev) TOTAL " +
					" from ( " +
					"           SELECT " +
					"               kpi.date as cdate, " +
					"               kpi.territory,  " +
					"               (kpi." + scol + " * (CAST(shar.country_rate AS FLOAT)/100)) AS rev  " +
					"               FROM kpi AS kpi  " +
					"               LEFT OUTER JOIN country_share AS shar " +
					"               ON kpi.territory = shar.territory  " +
					"               AND to_char(kpi.date,'YYYYMMDD') >= to_char(shar.start_date,'YYYYMMDD')  " +
					"               AND to_char(kpi.date,'YYYYMMDD') <= to_char(shar.end_date,'YYYYMMDD') " +
					"               where kpi.date >= ? and kpi.date <=  ? " +
					"               ORDER BY kpi.date, kpi.territory asc " +
					"		) a " +
					" group by ROLLUP(cdate)" +
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
		}
	} else {
		//나딕 매출 통계 프로세스 끝====================================================================================
		//일반 매출 통계 프로세스 시작==================================================================================
		//PERIOD SEARCH (1:DAILY, 2:WEEKLY, 3:MONTHLY)
		if period == "2" {

			//WEEKLY SEARCH
			if kind == "graph" {
				sql = " select cdate " +
					" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
					" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
					" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
					" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
					" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
					" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
					" ,sum(case when territory = 'VIETNAM' then rev else 0 end) VIETNAM" +
					" ,sum(rev) TOTAL " +
					" from ( " +
					"			select " +
					"				date(date_trunc('week', date))+6 as cdate " +
					"				, territory  " +
					"				, " + scol + " as rev  " +
					" 		from kpi where date(date_trunc('week', date))+6 >= ? and date <=  ? " +
					"			order by date, territory asc " +
					"		) a " +
					" group by ROLLUP(cdate)" +
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
		} else if period == "3" {//test
			if kind == "graph" {

				//MONTHLY SEARCH	//일자로 받은 데이터를 월까지만...
				from_substring := string(from[0:6])
				to_substring := string(to[0:6])

				sql = " select cdate " +
					" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
					" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
					" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
					" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
					" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
					" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
					" ,sum(case when territory = 'VIETNAM' then rev else 0 end) VIETNAM" +
					" ,sum(rev) TOTAL " +
					" from ( " +
					"			select " +
					"				to_char(date,'YYYY-MM') as cdate " +
					"				, territory  " +
					"				, " + scol + " as rev  " +
					" 		from kpi where to_char(date,'YYYYMM') >= ? and to_char(date,'YYYYMM') <=  ? " +
					"			order by date, territory asc " +
					"		) a " +
					" group by ROLLUP(cdate)" +
					" order by 1;"
				_, err = o.Raw(sql, from_substring, to_substring).QueryRows(&gListKpi)
			}

			if kind == "table" {
				sql = "SELECT territory, to_char(date,'YYYY-MM') as date, mcu_d as mcu, " +
					"avg_d as avg, uu_d as uu, nru_d as nru, " +
					"rev_d as rev, rev_ur_d as rev_ur, rev_pc_d as rev_pc, pu_d as pu, pur_d as pur, " +
					"arppu_d as arppu, dt, mrppu_d as mrppu, rev_t, rev_rate, bu " +
					"FROM kpi WHERE date >= ? and date <= ? "
				_, err = o.Raw(sql, from, to).QueryRows(&listKpi)
			}
		} else {
			if kind == "graph" {

				//DAILY SEARCH
				sql = " select cdate " +
					" ,sum(case when territory = 'KOREA' then rev else 0 end) KOREA" +
					" ,sum(case when territory = 'CHINA' then rev else 0 end) CHINA" +
					" ,sum(case when territory = 'JAPAN' then rev else 0 end) JAPAN" +
					" ,sum(case when territory = 'TAIWAN' then rev else 0 end) TAIWAN" +
					" ,sum(case when territory = 'NAMERICA' then rev else 0 end) NAMERICA" +
					" ,sum(case when territory = 'THAILAND' then rev else 0 end) THAILAND" +
					" ,sum(case when territory = 'VIETNAM' then rev else 0 end) VIETNAM" +
					" ,sum(rev) TOTAL " +
					" from ( " +
					"			select " +
					"				date as cdate " +
					"				, territory  " +
					"				, " + scol + " as rev  " +
					" 		from kpi where date >= ? and date <=  ? " +
					"			order by date, territory asc " +
					"		) a " +
					" group by ROLLUP(cdate)" +
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
		}
	}
	//일반 매출 통계 프로세스 끝====================================================================================

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
			" , gnru_d gnru" +
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
			" , sum(gnru_d) gnru" +
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
