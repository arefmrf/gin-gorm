package view

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"web/pkg/converters"
	"web/pkg/sessions"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	return data
}
