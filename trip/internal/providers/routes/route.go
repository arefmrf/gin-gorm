package routes

import (
	"github.com/gin-gonic/gin"
	publicHostRoutes "trip/internal/modules/public_host/routes"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	v1 := api.Group("/v1")
	publicHostRoutes.Routes(v1)
}
