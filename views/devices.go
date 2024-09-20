package views

import "net/http"

func DevicesHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Water District | Devices"
	return context
}
