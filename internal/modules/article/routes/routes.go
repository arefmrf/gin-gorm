package routes

import (
	"github.com/gin-gonic/gin"
	articleCtrl "web/internal/modules/article/controllers"
)

func Routes(router *gin.Engine) {
	articlesController := articleCtrl.New()
	router.GET("/article/:id", articlesController.Show)
}
