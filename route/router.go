package route

import (
	"github.com/gin-gonic/gin"
	"go_pro1/api"
)

func RouterCollect(r * gin.Engine) * gin.Engine  {
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/user/login", api.Login)
	return r
}