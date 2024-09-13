package views

import (
	"net/http"
	"strings"
)

func PageHandlers(w http.ResponseWriter, r *http.Request) {
	page := strings.TrimPrefix(r.URL.Path, "/")
	page = strings.TrimSuffix(page, "/")

	context := map[string]interface{}{}

	switch page {
	case "dashboard":
		context = DashboardHandler(w, r)
	}
	WebRender(w, r, page, context)
}
