package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/YoungsoonLee/backend_datainfra/models"
)

// ResourceController ...
type ResourceController struct {
	BaseController
}

// CreateResource ...
func (r *ResourceController) CreateResource() {
	var resource models.Resource
	body, _ := ioutil.ReadAll(r.Ctx.Request.Body)
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

// GetResourceAll ...
func (r *ResourceController) GetResourceAll() {
	// var resource models.Resource
	resource, err := models.GetResourceAll()
	if err != nil {
		r.ResponseError(libs.ErrDatabase, err)
	}

	r.ResponseSuccess("tabulator", &resource)
}
