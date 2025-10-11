package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/internal/modules/user/requests/auth"
	"web/pkg/app_errors"
	"web/pkg/converters"
	"web/pkg/sessions"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (controller *Controller) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	var request auth.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		app_errors.Init()
		app_errors.SetFromErrors(err)
		//c.JSON(http.StatusBadRequest, gin.H{"error": app_errors.Get()})
		sessions.Set(c, "errors", converters.MapToString(app_errors.Get()))
		c.JSON(http.StatusBadRequest, gin.H{"error": converters.MapToString(app_errors.Get())})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
