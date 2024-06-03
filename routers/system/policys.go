package system

import (
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/controllers/system"
)

func InitPolicyRouters(r *gin.RouterGroup) gin.IRoutes {
	{
		r.POST("/policy/createPolicy", system.AddCasbin)
		r.POST("/policy/deletePolicy", system.DelPolicy)
		r.GET("/policy/getPolicyList", system.ListPolicy)
	}
	return r
}
