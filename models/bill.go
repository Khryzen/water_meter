package models

import (
	"fmt"
	"time"

	"github.com/uadmin/uadmin"
	"golang.org/x/exp/rand"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Bill struct {
	uadmin.Model
	ReferenceNumber string
	Client          Client
	ClientID        uint
	AccountNumber   string
	BillingPeriod   BillingPeriod
	BillingPeriodID uint
	Amount          float64
	DueDate         time.Time
}

func (b *Bill) String() string {
	return b.AccountNumber + "- Php " + fmt.Sprint(b.Amount)
}

func (b *Bill) Save() {
	b.ReferenceNumber = GenerateRandomString(10)
	uadmin.Save(b)
}

func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
