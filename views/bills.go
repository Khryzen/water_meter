package views

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

func BillHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	bills := []models.Bill{}
	uadmin.All(&bills)

	context["Bills"] = bills
	context["Title"] = "Water District | Bills"
	return context
}
