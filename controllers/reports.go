package controllers

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

