package system

import (
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/controllers/system"
)

func InitUserRouters(r *gin.RouterGroup) gin.IRouter {
	{
		r.GET("/user/getUserInfo", system.GetUserInfo)
		r.GET("/user/getUserList", system.UserList)
		r.POST("/user/updateUser", system.UserUpdate)
		r.POST("/user/createUser", system.UserAdd)

	}
	return r
}
