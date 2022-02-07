package router

import (
	"github.com/gin-gonic/gin"
	."goschool/router/console"
)

func InitRouter(r *gin.Engine)  {
	InitConsoleRouter(r)
	//v1 := r.Group("/console")
	//{
	//	admin.initAdmin(v1)
	//	//news := v1.Group("/news")
	//	//{
	//	//	news.GET("/", service.SaveAdmin)
	//	//	//news.GET("/one", service.FindAdmin)
	//	//	//news.GET("/list", service.FindAdmins)
	//	//	//news.GET("/put", service.UpdateAdmin)
	//	//	//news.GET("/batchput", service.UpdateAdmins)
	//	//	//news.GET("/del", service.DeleteAdmin)
	//	//}
	//}
}
