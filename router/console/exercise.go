package console

import (
	"github.com/gin-gonic/gin"
	"go-website/service"

)

func ExerciseRouter(router *gin.RouterGroup) {
	r := router.Group("/exercise")
	r.POST("/", service.CreateExercise)
	r.GET("/:id", service.GetExercise)

}

