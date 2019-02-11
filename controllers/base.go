package controllers

import (
	"regexp"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

// ResponseError ...
func (b *BaseController) ResponseError(e *libs.ControllerError, err error) {
	// TODO: logging
	beego.Error(b.Ctx.Request.RequestURI, err)

	devInfo := ""
	if err != nil {
		devInfo = err.Error()
	}

	response := &models.ErrRespCode{
		Code:    e.Code,
		Message: e.Message,
		DevInfo: devInfo,
	}

	b.Ctx.Output.Status = e.Status
	b.Ctx.Output.JSON(response, true, true)

	b.StopRun()
}

// XsollaResponseError ...
func (b *BaseController) XsollaResponseError(e *libs.ControllerError) {
	// TODO: logging
	beego.Error(b.Ctx.Request.RequestURI, e.Message)

	eData := models.XRespDetailCode{
		Code:    e.Code,
		Message: e.Message,
	}

	response := &models.XRespCode{
		Error: eData,
	}

	b.Ctx.Output.Status = e.Status
	b.Ctx.Output.JSON(response, true, true)

	// TODO: logging
	b.StopRun()
}

// ValidDisplayname ...
func (b *BaseController) ValidDisplayname(displayname string) {

	if len(displayname) < 4 || len(displayname) > 16 {
		b.ResponseError(libs.ErrDisplayname, nil)
	}
}

// ValidID ...
func (b *BaseController) ValidID(id string) {

	if len(id) == 0 {
		b.ResponseError(libs.ErrIDAbsent, nil)
	}
}

// ValidEmail ...
func (b *BaseController) ValidEmail(email string) {
	valid := validation.Validation{}
	v := valid.Email(email, "Email")
	if !v.Ok {
		b.ResponseError(libs.ErrEmail, nil)
	}

	v = valid.MaxSize(email, 100, "Email")
	if !v.Ok {
		b.ResponseError(libs.ErrMaxEmail, nil)
	}
}

// ValidPassword ...
func (b *BaseController) ValidPassword(password string) {
	// 8 ~ 16 letters
	if len(password) < 8 || len(password) > 16 {
		b.ResponseError(libs.ErrPassword, nil)
	}

	valid := validation.Validation{}
	pattern := regexp.MustCompile("") //TODO: add regex for password

	v := valid.Match(password, pattern, "password")
	if !v.Ok {
		b.ResponseError(libs.ErrPassword, nil)
	}
}

/*
// ResponseHTTPError ...
func (b *BaseController) ResponseHTTPError(status int, code string, err error) {
	b.Ctx.Output.Status = status
	b.ResponseError(code, err)
}

// ResponseCommonError ...
func (b *BaseController) ResponseCommonError(e *libs.ControllerError) {
	beego.Error(fmt.Errorf(e.Message))
	b.ResponseHTTPError(e.Status, e.Code, fmt.Errorf(e.Message))
}

// ResponseServerError ...
func (b *BaseController) ResponseServerError(e *libs.ControllerError, err error) {
	beego.Error(err)
	b.ResponseHTTPError(e.Status, e.Code, fmt.Errorf(e.Message))
}

func loggingValidError(v *validation.Result) {
	beego.Error("key: ", v.Error.Key, ", value: ", v.Error.Value, ", message: ", v.Error.Message)
}
*/

// ResponseSuccess ...
func (b *BaseController) ResponseSuccess(key string, value interface{}) {
	b.Ctx.Output.Status = 200

	if key == "" {
		mresponse := &models.MrespCode{
			Code:    "ok",
			Message: "success",
			Data:    value,
		}

		b.Ctx.Output.JSON(mresponse, true, true)
	}

	if key == "tabulator" {
		b.Ctx.Output.JSON(value, true, true)
	}

	response := &models.RespCode{
		Code:    "ok",
		Message: "success",
		Data:    map[string]interface{}{},
	}

	response.Data[key] = value
	b.Ctx.Output.JSON(response, true, true)
}
