package system

import (
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/common/global"
	"new-go-edas-admin/controllers/system"
	"new-go-edas-admin/middles"
)

func InitBaseRouters(r *gin.RouterGroup) gin.IRouter {
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		global.TPLogger.Error("初始化JWT认证中间件失败：", err)
		panic(err)
	}
	{
		r.POST("/login", authMiddleware.LoginHandler)                           // 登录
		r.POST("/ldap/login", authMiddleware.LoginHandler)                      // ldap登录
		r.GET("/login/info", system.LoginUserInfo, authMiddleware.LoginHandler) // 登录用户详情
		r.POST("/logout", authMiddleware.LogoutHandler)                         // 退出
		r.POST("/refresh", authMiddleware.RefreshHandler)
	}
	return r
}
