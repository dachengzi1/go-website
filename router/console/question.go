package console

import (
	"github.com/gin-gonic/gin"
	"goschool/service"

)

func QuestionRouter(router *gin.RouterGroup) {
	r := router.Group("/quesion")
	r.POST("/", service.CreateQuestion)
	r.GET("/", service.SearchQuestions)
	r.PUT("/:id", service.UpdateQuestions)

}

