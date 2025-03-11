package models

import (
	"time"
)

type ResponseCfg struct {
	ResultCode int    `json:"result_code"`
	Message    string `json:"message"`
}

type HealthCfgRequest struct {
	Protocol string      `json:"protocol"`
	Response ResponseCfg `json:"response"`
	Interval int         `json:"interval"`
	IsActive bool        `json:"is_active"`
	Path     string      `json:"path"`
}

type HealthCfgResponse struct {
	Protocol   string      `json:"protocol"`
	Response   ResponseCfg `json:"response"`
	Interval   int         `json:"interval"`
	IsActive   bool        `json:"is_active"`
	Path       string      `json:"path"`
	LastUpdate time.Time   `json:"last_update"`
}
