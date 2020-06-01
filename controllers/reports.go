package controllers

import (
	"encoding/json"
	"io/ioutil"

	//"github.com/YoungsoonLee/backend_datainfra/libs"
	//"github.com/YoungsoonLee/backend_datainfra/models"

	"github.com/astaxie/beego/logs"
	"github.com/ndcinfra/backend_datainfra/libs"
)

// KpiController ...
type ReportsController struct {
	BaseController
}

// Client Data
type InClientData struct {
	clCountry  string `json:"clcoun"` // client country
	clVersion  string `json:"clver"`  // client version
	clServerIP string `json:"clsip"`  // connect server ip
	clSVersion string `json:"clsver"` // connect server version
	clFile     string `json:"clfile"` // Login_address.lad file Y/N
	clConnect  string `json:"clcon"`  // connect server success Y/N
}

// Server Data
type InServerData struct {
	sCountry string `json:"clver"`  // server country
	sBinary  string `json:"clcoun"` // server binary
	sVersion string `json:"clver"`  // server version
	sPuIP    string `json:"clsip"`  // server public ip
	sPrIP    string `json:"clcon"`  // server private ip
	sFarmIP  string `json:"clsver"` // farm server ip
}

// Reports 수집용. Client
func (k *ReportsController) GetClient() {
	var inCDate InClientData

	// input check.. 값이 하나라도 없으면 에러
	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inCDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// Log Write
	logs.Info(&inCDate)
}

// Reports 수집용. server
func (k *ReportsController) GetServer() {
	var inSDate InServerData

	// input check.. 값이 하나라도 없으면 에러
	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inSDate)
	if err != nil {
		k.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// Log Write
	logs.Info(&inSDate)
}
