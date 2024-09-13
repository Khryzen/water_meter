package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
)

type Bill struct {
	uadmin.Model
	ReferenceNumber string
	AccountNumber   string
	BillingPeriod   BillingPeriod
	BillingPeriodID uint
	Reading         Reading
	ReadingID       uint
	Amount          float64
	DueDate         time.Time
}

func (b *Bill) String() string {
	return b.AccountNumber + "- Php " + fmt.Sprint(b.Amount)
}
