package routes

import (
	"github.com/gin-gonic/gin"
	homeRoutes "web/internal/modules/home/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
}
