package log

import (
	"time"
)

type Log struct {
	Filename string    `json:"filename"`
	Artist   string    `json:"artist"`
	Song     string    `json:"song"`
	Time     time.Time `json:"time"`
	Hour     string    `json:"hour"`
	Day      string    `json:"day"`
	Month    string    `json:"month"`
	Year     string    `json:"year"`
}
