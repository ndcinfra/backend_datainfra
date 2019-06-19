package controllers

import "time"

// KpiController ...
type SystemController struct {
	BaseController
}

func (s *SystemController) CheckHealthy() {
	k.ResponseSuccess("", time.Now())
}
