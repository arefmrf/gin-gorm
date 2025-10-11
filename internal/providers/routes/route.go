package routes

import (
	"github.com/gin-gonic/gin"
	articleRoutes "web/internal/modules/article/routes"
	homeRoutes "web/internal/modules/home/routes"
	usereRoutes "web/internal/modules/user/routes"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	usereRoutes.Routes(router)
}
