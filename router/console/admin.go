package console

import (
	"github.com/gin-gonic/gin"
	"go-website/service"
)


func AdminRouter(router *gin.RouterGroup) {
	r := router.Group("/admin")

	r.POST("/", service.SaveAdmin)
	r.GET("/", service.FindAdmins)
	r.GET("/:id", service.FindAdmin)
	r.PUT("/put", service.UpdateAdmin)
	r.PUT("/batch", service.UpdateAdmins)
	r.DELETE("/del", service.DeleteAdmin)

}
