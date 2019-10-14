package controllers

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	//"github.com/YoungsoonLee/backend_datainfra/libs"
	//"github.com/YoungsoonLee/backend_datainfra/models"

	"github.com/ndcinfra/backend_datainfra/libs"
	"github.com/ndcinfra/backend_datainfra/models"
)

// ResourceController ...
type ResourceController struct {
	BaseController
}

// CreateResource ...
func (r *ResourceController) CreateResource() {
	var resource models.Resource
	body, _ := ioutil.ReadAll(r.Ctx.Request.Body)
	//fmt.Println(body)

	err := json.Unmarshal(body, &resource)
	if err != nil {
		r.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// save to db
	ID, err := models.AddResource(resource)
	if err != nil || ID == -1 {
		r.ResponseError(libs.ErrDatabase, err)
	}

	// auto login
	resource.ID = ID
	r.ResponseSuccess("", &resource)
}

// UpdateResource ...
func (r *ResourceController) UpdateResource() {
	id := r.GetString(":id")
	if len(id) == 0 {
		r.ResponseError(libs.ErrResourceIDAbsent, nil)
	}

	intID, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		r.ResponseError(libs.ErrConvert, err)
	}

	var resource models.Resource
	body, _ := ioutil.ReadAll(r.Ctx.Request.Body)

	err = json.Unmarshal(body, &resource)
	if err != nil {
		r.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	resource.ID = intID
	//fmt.Println(resource)

	_, err = models.UpdateResource(resource)
	if err != nil {
		r.ResponseError(libs.ErrDatabase, err)
	}

	r.ResponseSuccess("", &resource)

}

// DeleteResource ...
func (r *ResourceController) DeleteResource() {
	id := r.GetString(":id")
	if len(id) == 0 {
		r.ResponseError(libs.ErrResourceIDAbsent, nil)
	}

	intID, err := strconv.ParseInt(id, 10, 64)

	err = models.DeleteResource(intID)
	if err != nil {
		r.ResponseError(libs.ErrDatabase, err)
	}

	r.ResponseSuccess("", "")

}

// GetResourceAll ...
func (r *ResourceController) GetResources() {
	// var resource models.Resource
	resource, err := models.GetResources()
	if err != nil {
		r.ResponseError(libs.ErrDatabase, err)
	}

	r.ResponseSuccess("tabulator", &resource)
}

func (r *ResourceController) GetResourceDetail() {

	id := r.GetString(":id")

	if len(id) == 0 {
		r.ResponseError(libs.ErrResourceIDAbsent, nil)
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		r.ResponseError(libs.ErrConvert, err)
	}

	resource, err := models.GetResourceDetail(intID)
	if err != nil {
		r.ResponseError(libs.ErrDatabase, err)
	}

	r.ResponseSuccess("", &resource)

}
