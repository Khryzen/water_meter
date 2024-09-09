package models

import "github.com/uadmin/uadmin"

type Reading struct {
	uadmin.Model
	Device          Device
	DeviceID        uint
	BillingPeriod   BillingPeriod
	BillingPeriodID uint
	Beginning       float64
	Ending          float64
	Consumption     float64
}
