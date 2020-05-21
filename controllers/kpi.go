package controllers

import (
	"encoding/json"
	"io/ioutil"

	//"github.com/YoungsoonLee/backend_datainfra/libs"
	//"github.com/YoungsoonLee/backend_datainfra/models"

	"github.com/ndcinfra/backend_datainfra/libs"
	"github.com/ndcinfra/backend_datainfra/models"
)

// KpiController ...
type KpiController struct {
	BaseController
}

type InputDate struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Country      string `json:"country"`
	Kind         string `json:"kind"`         // graph, table
	Radio        string `json:"radio"`        // uu, mcu ...
	KindCalendar string `json:"kindCalendar"` // day, week, month
	Period       string `json:"period"`       // day, week, month
}

// GetKPI ...
// Dashboard 용
func (k *KpiController) GetKPI() {
	var inputDate InputDate

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inputDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	var kpi models.Kpi

	listKpi, gListKpi, err := kpi.GetKPI(inputDate.From, inputDate.To, inputDate.Country, inputDate.Kind, inputDate.Radio)
	if err != nil {
		k.ResponseError(libs.ErrDatabase, err)
	}

	if inputDate.Kind == "graph" {
		k.ResponseSuccess("", gListKpi)
	}
	k.ResponseSuccess("", listKpi)

}

// Dashboard 용
func (k *KpiController) GetNewKPI() {
	var inputDate InputDate

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inputDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	var kpi models.Kpi

	listKpi, gListKpi, err := kpi.GetNewKPI(inputDate.From, inputDate.To, inputDate.Country, inputDate.Kind, inputDate.Radio, inputDate.Period)
	if err != nil {
		k.ResponseError(libs.ErrDatabase, err)
	}

	if inputDate.Kind == "graph" {
		k.ResponseSuccess("", gListKpi)
	}
	k.ResponseSuccess("", listKpi)

}

// GetUserKPI
// 유저 통계 용
func (k *KpiController) GetUserKPI() {
	var inputDate InputDate

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inputDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	var kpi models.UserKPI

	listKpi, err := kpi.GetUserKPI(inputDate.From, inputDate.To, inputDate.Country, inputDate.Kind, inputDate.Radio, inputDate.KindCalendar)
	if err != nil {
		k.ResponseError(libs.ErrDatabase, err)
	}

	/*
		if inputDate.Kind == "graph" {
			k.ResponseSuccess("", gListKpi)
		}
	*/

	k.ResponseSuccess("", listKpi)
}

// GetSaleKPI
// 매출 통계 용
func (k *KpiController) GetSaleKPI() {
	var inputDate InputDate

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inputDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	var kpi models.SaleKPI

	listKpi, err := kpi.GetSaleKPI(inputDate.From, inputDate.To, inputDate.Country, inputDate.Kind, inputDate.Radio, inputDate.KindCalendar)
	if err != nil {
		k.ResponseError(libs.ErrDatabase, err)
	}

	/*
		if inputDate.Kind == "graph" {
			k.ResponseSuccess("", gListKpi)
		}
	*/

	// logs.Info(listKpi)

	k.ResponseSuccess("", listKpi)
}

// 기타
// 아이템 Top 50
func (k *KpiController) GetSaleItemKPI() {
	var inputDate InputDate

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inputDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	var kpi models.SaleItemKPI

	listKpi, err := kpi.GetSaleItemKPI(inputDate.From, inputDate.To, inputDate.Country, inputDate.Kind, inputDate.Radio, inputDate.KindCalendar)
	if err != nil {
		k.ResponseError(libs.ErrDatabase, err)
	}

	//logs.Info(listKpi)
	k.ResponseSuccess("", listKpi)
}
