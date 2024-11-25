package models

import (
	"encoding/json"
	"time"
)

type HealthCfg struct {
	ID         string          `json:"id"`       
	Protocol   string          `json:"protocol"`
	Response   json.RawMessage `json:"response"`
	Interval   int             `json:"interval"`
	IsActive   bool            `json:"is_active"`
	LastUpdate time.Time       `json:"last_update"`
}
