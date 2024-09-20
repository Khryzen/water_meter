package views

import "net/http"

func ClientsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Water District | Clients"
	return context
}
