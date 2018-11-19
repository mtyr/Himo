package model

import "time"

type Alarm struct {
	SeetID         string    `json:"seet_id"`
	Time           time.Time `json:"time"`
	IsLongInterval bool      `json:"is_long_interval"`
}
