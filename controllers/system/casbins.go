package system

import (
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/common/global"
	mod "new-go-edas-admin/models/system"
	service "new-go-edas-admin/service/system"
)

func AddCasbin(ctx *gin.Context) {
	params := new(struct {
		Policy []*mod.CasbinPolicy `json:"policy"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		global.TPLogger.Error("添加授权参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := service.NewCasbinInterface().AddPolicy(params.Policy)
	if err != nil {
		global.TPLogger.Error("添加授权失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
	}
	if err = global.CasbinEnforcer.LoadPolicy(); err != nil {
		global.TPLogger.Error("加载权限失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", "添加权限成功")
}

// 删除授权
func DelPolicy(ctx *gin.Context) {
	params := new(struct {
		Policy []*mod.CasbinPolicy `json:"policy"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		global.TPLogger.Error("删除授权参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := service.NewCasbinInterface().DelPolicy(params.Policy)
	if err != nil {
		global.TPLogger.Error("删除授权失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	if err = global.CasbinEnforcer.LoadPolicy(); err != nil {
		global.TPLogger.Error("加载权限失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
	}
	global.ReturnContext(ctx).Successful("success", "删除权限成功")
}

// 获取权限列表
func ListPolicy(ctx *gin.Context) {
	params := new(struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		global.TPLogger.Error("获取权限列表参数校验失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	data := service.NewCasbinInterface().ListPolicy(params.Page, params.Limit)
	global.ReturnContext(ctx).Successful("success", data)

}
