package models

import (
	"fmt"
)

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
	sCountry string `json:"clcoun"` // server country
	sBinary  string `json:"clbin"`  // server binary
	sVersion string `json:"clver"`  // server version
	sPuIP    string `json:"clpuip"` // server public ip
	sPrIP    string `json:"clprip"` // server private ip
	sFarmIP  string `json:"clfip"`  // farm server ip
}

// GetUserKPI ...
func (k *InClientData) GetClientReport(clCountry, clVersion, clServerIP, clSVersion, clFile, clConnect string) {

	fmt.Println("client input: ", clCountry, clVersion, clServerIP, clSVersion, clFile, clConnect)
}

func (k *InServerData) GetServerReport(sCountry, sBinary, sVersion, sPuIP, sPrIP, sFarmIP string) {
	fmt.Println("server input: ", sCountry, sBinary, sVersion, sPuIP, sPrIP, sFarmIP)
}
