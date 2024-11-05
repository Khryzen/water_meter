package views

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	devices := models.Device{}
	bill := models.Bill{}

	context["Bill"] = uadmin.Count(&bill, "paid = ?", false)
	context["Paid"] = uadmin.Count(&bill, "paid = ?", true)
	context["Device"] = uadmin.Count(&devices, "active = ?", true)
	context["Title"] = "Water District | Dashboard"
	return context
}
