package controllers

import (
	"net/http"
	"strconv"
	"trip/internal/modules/public_host/requests"
	PublicHostService "trip/internal/modules/public_host/services"
	"trip/pkg/pagination"

	"github.com/gin-gonic/gin"
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
	var req requests.PublicHostListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "er3000",
		})
		return
	}
	
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	p := pagination.New(page, limit)

	hosts := controller.publicHostService.GetPublicHosts(p)
	c.JSON(http.StatusOK, gin.H{"hosts": hosts})
}
