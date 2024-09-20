package views

import "net/http"

func PaymentsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Water District | Payments"
	return context
}
