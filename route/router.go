package route

import (
	"github.com/gin-gonic/gin"
	"go_pro1/api"
	"go_pro1/middleware"
)

func RouterCollect(r * gin.Engine) * gin.Engine  {
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/userinfo", middleware.AuthMiddleware(), api.Userinfo)

	return r
}