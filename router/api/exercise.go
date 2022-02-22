package api

import (
	"github.com/gin-gonic/gin"
	"go-website/middleware"
	"go-website/service"
)

func ExerciseRouter(router *gin.RouterGroup) {
	r := router.Group("exercise")
	r.Use(middleware.ApiAuth())

	r.POST("/save", service.SaveUserExercise)
	r.GET("/:id", service.GetUserExercise)

}
