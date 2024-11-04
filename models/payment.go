package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Payment struct {
	uadmin.Model
	ReferenceNumber string
	Bill            Bill
	BillID          uint
	Amount          float64
	PaymentDate     *time.Time
}
