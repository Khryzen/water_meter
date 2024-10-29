package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Reading struct {
	uadmin.Model
	DateString  string
	Date        time.Time
	Device      Device
	DeviceID    uint
	Beginning   float64
	Ending      float64
	Consumption float64
	Client      Client
	ClientID    uint
}

func (r *Reading) Save() {
	layout := "01022006" // Parsing layout for MMDDYYYY

	// Parse the date string
	parsedDate, err := time.Parse(layout, r.DateString)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Error parsing date:", err)
		return
	}
	r.Date = parsedDate
	uadmin.Save(r)
}
