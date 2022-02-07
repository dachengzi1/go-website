package console

import (
	"github.com/gin-gonic/gin"
	"go-website/service"

)

func QuestionRouter(router *gin.RouterGroup) {
	r := router.Group("/quesion")
	r.POST("/", service.CreateQuestion)
	r.GET("/", service.SearchQuestions)
	r.PUT("/:id", service.UpdateQuestions)

}

