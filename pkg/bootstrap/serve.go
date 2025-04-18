package bootstrap

import (
	"web/pkg/config"
	"web/pkg/database"
	"web/pkg/html"
	"web/pkg/routing"
	"web/pkg/static"
)

func Serve() {
	config.Set()
	database.Connect()
	routing.Init()

	static.LoadStatic(routing.GetRouter())
	html.LoadHTML(routing.GetRouter())
	routing.RegisterRoutes()

	routing.Serve()
}
