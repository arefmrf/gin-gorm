package routes

import (
	"github.com/gin-gonic/gin"
	publicHostCtrl "trip/internal/modules/public_host/controllers"
)

func Routes(r *gin.RouterGroup) {
	publicHostController := publicHostCtrl.New()
	hosts := r.Group("public/hosts")
	hosts.GET("/general", publicHostController.List)
}
