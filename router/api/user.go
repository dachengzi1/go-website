package api

import (
	"github.com/gin-gonic/gin"
	"go-website/service"
)

func InitUserRouter(router *gin.RouterGroup) {
	r := router.Group("/user")
	r.POST("/", service.SaveUser)
}
