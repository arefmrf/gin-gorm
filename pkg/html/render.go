package html

import (
	"github.com/gin-gonic/gin"
	"web/internal/providers/view"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = view.WithGlobalData(c, data)
	c.HTML(code, name, data)
}
