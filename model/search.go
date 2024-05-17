package model

import "time"

type SearchCriteria struct {
	Make      string
	Model     string
	StartTime time.Time
	EndTime   time.Time
	MinPrice  float64
	MaxPrice  float64
}
