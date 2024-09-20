package views

import "net/http"

func BillingPeriodHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Water District | Billing Period"
	return context
}
