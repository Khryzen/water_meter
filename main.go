package main

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/mbdeguzman/water_district/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Bill{},
		models.BillingPeriod{},
		models.Client{},
		models.Device{},
		models.Payment{},
		models.Reading{},
	)

	http.HandleFunc("/", uadmin.Handler(views.PageHandlers))
	uadmin.Port = 8080
	uadmin.RootURL = "/uadmin/"

	uadmin.StartServer()
}
