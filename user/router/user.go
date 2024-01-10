package router

import (
	"user-services/handle"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, h handle.UserHandle) {
	r.POST("/sign-up", h.SignUp)
	r.POST("/sign-in", h.SignIn)
}
