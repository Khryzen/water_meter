package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type BillingPeriod struct {
	uadmin.Model
	Month      int
	Year       int
	Multiplier float64
	DueDate    time.Time
}

// Create a save function that will generate the bills
func (b *BillingPeriod) Save() {
	uadmin.Save(b)

	clients := []Client{}
	uadmin.Filter(&clients, "active = ?", true)

	for i := range clients {
		readings := []Reading{}
		uadmin.Filter(&readings, "month = ? AND year = ? AND client_id = ?", b.Month, b.Year, clients[i].ID)
		bill := Bill{}
		bill.ClientID = clients[i].ID
		bill.BillingPeriodID = b.ID
		consumption := 0.00
		for j := range readings {
			consumption += readings[j].Consumption
		}
		bill.Amount = b.Multiplier * consumption
		bill.Save()
	}
}
