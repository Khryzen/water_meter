package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type BillingPeriod struct {
	uadmin.Model
	Month      int
	Multiplier float64
	DueDate    time.Time
}
