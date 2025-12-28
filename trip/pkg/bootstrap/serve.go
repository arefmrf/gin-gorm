package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"trip/pkg/config"
	"trip/pkg/database"
	"trip/pkg/routing"
)

type Server struct {
}

func Run() {
	log.Printf("🚀 Server running at http://localhost\n")
	config.Set()
	configs := config.Get()
	database.Connect()

	routing.Init()
	routing.RegisterRoutes()

	addr := fmt.Sprintf(":%s", configs.Server.Port)
	log.Printf("🚀 Server running at http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, routing.GetRouter()); err != nil {
		log.Fatal(err)
	}
}
