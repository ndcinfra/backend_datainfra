package controllers

import "time"

// KpiController ...
type SystemController struct {
	BaseController
}

func (s *SystemController) CheckHealthy() {
	s.ResponseSuccess("", time.Now())
}
