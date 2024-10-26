package views

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

func ClientsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}
	devices := []models.Device{}
	uadmin.Filter(&devices, "assigned = ?", false)

	clients := []models.Client{}
	uadmin.All(&clients)
	for i := range clients {
		uadmin.Preload(&clients[i])
	}

	context["Clients"] = clients
	context["Devices"] = devices
	context["Title"] = "Water District | Clients"
	return context
}
