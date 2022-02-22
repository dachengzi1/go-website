package api

import (
	"github.com/gin-gonic/gin"
)

func InitApiRouter(r *gin.Engine) {

	v1 := r.Group("/api")
	{

		ExerciseRouter(v1)
		InitUserRouter(v1)

	}
}
