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
	default:
		context = DashboardHandler(w, r)
		page = "dashboard"
	}
	WebRender(w, r, page, context)
}
