package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"
)

// KpiController ...
type KpiController struct {
	BaseController
}

type InputDate struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Country string `json:"country"`
	Kind    string `json:"kind"` // graph, table
}

// CreateResource ...
func (k *KpiController) GetKPI() {
	var inputDate InputDate

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inputDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	var kpi models.Kpi

	listKpi, gListKpi, err := kpi.GetKPI(inputDate.From, inputDate.To, inputDate.Country, inputDate.Kind)
	if err != nil {
		k.ResponseError(libs.ErrDatabase, err)
	}

	if inputDate.Kind == "graph" {
		k.ResponseSuccess("", gListKpi)
	}
	k.ResponseSuccess("", listKpi)

}
