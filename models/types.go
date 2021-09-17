package models

import "time"

type TimeResponse struct {
	Zone   string    `json:"zone"`
	Time   time.Time `json:"time"`
	Offset string    `json:"offset"`
}
