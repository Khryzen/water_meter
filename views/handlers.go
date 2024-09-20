package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func PageHandlers(w http.ResponseWriter, r *http.Request) {
	page := strings.TrimPrefix(r.URL.Path, "/")
	page = strings.TrimSuffix(page, "/")

	uadmin.Trail(uadmin.DEBUG, "page: %s", page)

	context := map[string]interface{}{}

	switch page {
	case "dashboard":
		context = DashboardHandler(w, r)
		page = "dashboard"
	case "devices":
		context = DevicesHandler(w, r)
		page = "devices"
	case "clients":
		context = ClientsHandler(w, r)
		page = "clients"
	case "payments":
		context = PaymentsHandler(w, r)
		page = "payments"
	case "bill":
		context = BillHandler(w, r)
		page = "bill"
	case "billing_period":
		context = BillingPeriodHandler(w, r)
		page = "billing_period"
	case "readings":
		context = ReadingsHandler(w, r)
		page = "readings"
	default:
		context = DashboardHandler(w, r)
		page = "dashboard"
	}
	WebRender(w, r, page, context)
}
