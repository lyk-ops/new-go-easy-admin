package system

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitBaseRouters(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRouter {
	//authMiddleware, err := middles.InitAuth()
	//if err != nil {
	//	global.TPLogger.Error("初始化JWT认证中间件失败：", err)
	//	panic(err)
	//}
	{
		//r.POST("/login", authMiddleware.LoginHandler) // 登录
		//r.POST("/ldap/login", authMiddleware.LoginHandler)                      // ldap登录
		//r.GET("/login/info", system.LoginUserInfo, authMiddleware.LoginHandler) // 登录用户详情
		//r.POST("/logout", authMiddleware.LogoutHandler)                         // 退出
		//r.POST("/refresh", authMiddleware.RefreshHandler)
		r.POST("/login", authMiddleware.LoginHandler) // 登录
	}
	return r
}
