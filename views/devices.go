package views

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

func DevicesHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	devices := []models.Device{}
	uadmin.All(&devices)
	context["Devices"] = devices
	context["Title"] = "Water District | Devices"
	return context
}
