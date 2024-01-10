package router

import (
	"user-services/di"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, d *di.ServiceDI) {
	UserRouter(r, d.UserHandle)
}
