package main

import (
	"time"
)

type responseCfg struct {
	resultCode int    `json:"result code"`
	message    string `json:"message"`
}

type healthCfg struct {
	ID         string      `json:"id"`
	Protocol   string      `json:"protocol"`
	Response   responseCfg `json:"response"`
	Interval   int         `json:"interval"`
	IsActive   bool        `json:"is_active"`
	LastUpdate time.Time   `json:"last_update"`
}
