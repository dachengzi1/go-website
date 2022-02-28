package api

import (
	"github.com/gin-gonic/gin"
	"go-website/middleware"
	"go-website/service"
)

func InitPubRouter(group *gin.RouterGroup) {
	r := group.Group("/pub")
	r.Use(middleware.ApiAuth())

	{
		r.GET("/config", service.GetConfig)
	}
}
