package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	PublicHostService "trip/internal/modules/public_host/services"
	"trip/pkg/pagination"
)

type Controller struct {
	publicHostService PublicHostService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		publicHostService: PublicHostService.New(),
	}
}

func (controller *Controller) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	p := pagination.New(page, limit)

	hosts := controller.publicHostService.GetPublicHosts(p)
	c.JSON(http.StatusOK, gin.H{"hosts": hosts})
}
