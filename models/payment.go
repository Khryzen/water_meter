package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Payment struct {
	uadmin.Model
	ReferenceNumber string
	Amount          float64
	PaymentDate     *time.Time
}
