package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func WebRender(w http.ResponseWriter, r *http.Request, tpl string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/web/base.html")

	path := "./templates/web/" + tpl + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
