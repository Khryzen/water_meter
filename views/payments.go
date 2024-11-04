package views

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

func PaymentsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	payments := []models.Payment{}
	uadmin.All(&payments)
	for i := range payments {
		uadmin.Preload(&payments[i])
	}

	bills := []models.Bill{}
	uadmin.Filter(&bills, "paid = ?", false)

	context["Payments"] = payments
	context["Bills"] = bills
	context["Title"] = "Water District | Payments"
	return context
}
