package system

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/controllers/system"
)

func InitUserRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRouter {
	{
		r.POST("/user/logout", authMiddleware.LogoutHandler)   // 退出
		r.POST("/user/refresh", authMiddleware.RefreshHandler) // 刷新令牌
		r.GET("/user/getUserInfo", system.GetUserInfo)
		r.GET("/user/getUserList", system.UserList)
		r.POST("/user/updateUser", system.UserUpdate)
		r.POST("/user/createUser", system.UserAdd)

	}
	return r
}
