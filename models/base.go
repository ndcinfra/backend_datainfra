package models

// one return value
type RespCode struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	DevInfo string                 `json:"devinfo"`
	Data    map[string]interface{} `json:"data"`
}

// one return value
type ErrRespCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	DevInfo string `json:"devinfo"`
}

// multi return value
type MrespCode struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	DevInfo string      `json:"devinfo"`
	Data    interface{} `json:"data"`
}

// multi return value for BI
type MrespCodeBI struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	DevInfo string      `json:"devinfo"`
	Data    interface{} `json:"data"`
	Data2   interface{} `json:"data2"`
}

type XRespDetailCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type XRespCode struct {
	Error XRespDetailCode `json:"error"`
}
