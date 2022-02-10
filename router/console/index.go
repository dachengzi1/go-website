package console

import (
	"github.com/gin-gonic/gin"
)

func InitConsoleRouter(r *gin.Engine) {

	v1 := r.Group("/console")
	{
		AdminRouter(v1)
		QuestionRouter(v1)
		ExerciseRouter(v1)

	}

	//v1 := r.Group("/console")
	//{
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
