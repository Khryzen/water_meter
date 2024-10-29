package models

import (
	"strconv"
	"time"

	"github.com/uadmin/uadmin"
)

type Reading struct {
	uadmin.Model
	DateString  string
	Month       int
	Day         int
	Year        int
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
	layout := "01022006"

	month, _ := strconv.ParseInt(r.DateString[0:2], 10, 64)
	day, _ := strconv.ParseInt(r.DateString[2:4], 10, 64)
	year, _ := strconv.ParseInt(r.DateString[4:8], 10, 64)

	r.Month = int(month)
	r.Day = int(day)
	r.Year = int(year)

	// Parse the date string
	parsedDate, err := time.Parse(layout, r.DateString)
	if err != nil {
		uadmin.Trail(uadmin.ERROR, "Error parsing date:", err)
		return
	}
	r.Date = parsedDate
	uadmin.Save(r)
}
