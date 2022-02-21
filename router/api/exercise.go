package api

import (
	"github.com/gin-gonic/gin"
	"go-website/service"
)

func ExerciseRouter(router *gin.RouterGroup) {
	r := router.Group("exercise")

	r.POST("/save", service.SaveUserExercise)
	r.GET("/:id", service.GetUserExercise)

}
