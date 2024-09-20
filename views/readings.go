package views

import "net/http"

func ReadingsHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	context := map[string]interface{}{}

	context["Title"] = "Water District | Readings"
	return context
}
