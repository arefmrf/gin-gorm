package routes

import (
	"github.com/gin-gonic/gin"
	homeCtrl "web/internal/modules/home/controllers"
)

func Routes(router *gin.Engine) {
	homeController := homeCtrl.New()
	router.GET("/", homeController.Index)
	
	//router.GET("/", func(c *gin.Context) {
	//c.HTML(http.StatusOK, "modules/home/html/home", gin.H{
	//	"title":    "Home page",
	//	"APP_NAME": viper.Get("App.Name"),
	//})

	//html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
	//	"title": "Home page",
	//})
	//})

	//router.GET("/about", func(c *gin.Context) {
	//	html.Render(c, http.StatusOK, "modules/home/html/about", gin.H{
	//		"title": "About page",
	//	})
	//})
}
