package domain

import "time"

type (
	Rule struct {
		Qty       uint
		RangeTime time.Duration
	}

	GroupName string

	Notification struct {
		UserID      string
		CreatedDate time.Time
	}
)
