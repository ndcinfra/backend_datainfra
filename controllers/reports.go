package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ReportsController struct {
	BaseController
}

type InClientData struct {
	ClCountry  string `json:"clcoun"`
	ClVersion  string `json:"clver"`
	ClServerIP string `json:"clsip"`
	ClSVersion string `json:"clsver"`
	ClFile     string `json:"clfile"`
	ClConnect  string `json:"clcon"`
}

type InServerData struct {
	SCountry   string `json:"scoun"`
	SBinary    string `json:"sbin"`
	SVersion   string `json:"sver"`
	SPublicIP  string `json:"spu"`
	SPraviteIP string `json:"spr"`
	SFarmIP    string `json:"sfar"`
}

func (k *ReportsController) GetClient() {

	var inclientdata InClientData

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inclientdata)

	if err != nil {
		fmt.Println("client error: ", err)
		return
	}

	// Success Log Write
	k.writeLog("info", "Client1 : ", inclientdata)
}

func (k *ReportsController) GetServer() {
	var inserverdata InServerData

	body, _ := ioutil.ReadAll(k.Ctx.Request.Body)
	err := json.Unmarshal(body, &inserverdata)

	if err != nil {
		// Input Error
		fmt.Println("server error: ", err)
		return
	}

	// Success Log Write
	k.writeLog("info", "Server : ", inserverdata)
}
