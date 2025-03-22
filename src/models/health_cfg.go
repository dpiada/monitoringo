package models

import (
	"time"
)

type ResponseCfg struct {
	ResultCode int    `json:"result_code"`
	Message    string `json:"message"`
}

type HealthCfgRequest struct {
	Method   string      `json:"Method"`
	Response ResponseCfg `json:"response"`
	Interval int         `json:"interval"`
	IsActive bool        `json:"is_active"`
	Path     string      `json:"path"`
}

type HealthCfgResponse struct {
	Method     string      `json:"Method"`
	Response   ResponseCfg `json:"response"`
	Interval   int         `json:"interval"`
	IsActive   bool        `json:"is_active"`
	Path       string      `json:"path"`
	LastUpdate time.Time   `json:"last_update"`
}
