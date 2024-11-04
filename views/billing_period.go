package views

import (
	"net/http"
	"time"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

type Period struct {
	Month      string
	Year       int
	Multiplier float64
	DueDate    time.Time
}

func BillingPeriodHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	billing_periods := []models.BillingPeriod{}
	uadmin.All(&billing_periods)

	periods := []Period{}

	for i := range billing_periods {
		periods = append(periods, Period{
			Month:      ToMonthString(billing_periods[i].Month),
			Year:       billing_periods[i].Year,
			Multiplier: billing_periods[i].Multiplier,
			DueDate:    billing_periods[i].DueDate,
		})
	}

	context["BillingPeriods"] = billing_periods
	context["Title"] = "Water District | Billing Period"
	return context
}

func ToMonthString(month int) string {
	var monthStr string
	if month == 1 {
		monthStr = "January"
	} else if month == 2 {
		monthStr = "February"
	} else if month == 3 {
		monthStr = "March"
	} else if month == 4 {
		monthStr = "April"
	} else if month == 5 {
		monthStr = "May"
	} else if month == 6 {
		monthStr = "June"
	} else if month == 7 {
		monthStr = "July"
	} else if month == 8 {
		monthStr = "August"
	} else if month == 9 {
		monthStr = "September"
	} else if month == 10 {
		monthStr = "October"
	} else if month == 11 {
		monthStr = "November"
	} else if month == 12 {
		monthStr = "December"
	} else {
		monthStr = "Invalid"
	}

	return monthStr
}
