package models

import "github.com/uadmin/uadmin"

type Reading struct {
	uadmin.Model
	DateString  string
	Device      Device
	DeviceID    uint
	Beginning   float64
	Ending      float64
	Consumption float64
	Client      Client
	ClientID    uint
}
