package controllers

<<<<<<< HEAD
import {
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego/logs"
	"github.com/ndcinfra/backend_datainfra/libs"
}

type RepotsController struct {
	BaseController
}

type InClientData struct {
	clCountry 	string `json:"clcoun"`
	clVersion 	string `json:"clver"`
	clServerIP 	string `json:"clsip"`
	clVersion 	string `json:"clsver"`
	clFile 		string `json:"clfile"`
	clConnect 	string `json:"clcon"`
}

type InServerData struct {
	sCountry 	string `json:"scoun"`
	sBinary		string `json:"sbin"`
	sVersion	string `json:"sver"`
	sPublicIP	string `json:"spu"`
	sPraviteIP	string `json:"spr"`
	sFarmIP		string `json:"sfar"`
}

func (k *ReportsController) Getcli() {
	var inCData InClientData

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inCData)

	if  err != nil {
		// Input Error
	}

	// Success Log Write
	logs.Info(&inCData)
}

func (k *ReportsController) GetSer() {
	var inSData InServerData

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inSData)

	if  err != nil {
		// Input Error
	}

	// Success Log Write
	logs.Info(&inSData)
}

=======
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
>>>>>>> ae9d5536a882204515427b087385a68ca4352d25
