package router

import (
	"github.com/gin-gonic/gin"
	"go-website/router/api"
	"go-website/router/console"
)

func InitRouter(r *gin.Engine) {
	console.InitConsoleRouter(r)
	api.InitApiRouter(r)
}
