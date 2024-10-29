package views

import (
	"net/http"

	"github.com/mbdeguzman/water_district/models"
	"github.com/uadmin/uadmin"
)

func ReadingsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	readings := []models.Reading{}
	uadmin.All(&readings)

	for i := range readings {
		uadmin.Preload(&readings[i])
	}

	context["Readings"] = readings
	context["Title"] = "Water District | Readings"
	return context
}
