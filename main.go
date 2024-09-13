package main

import (
	"net/http"

	"github.com/mbdeguzman/water_district/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register()

	http.HandleFunc("/", uadmin.Handler(views.PageHandlers))
	uadmin.Port = 8080
	uadmin.RootURL = "/uadmin/"

	uadmin.StartServer()
}
