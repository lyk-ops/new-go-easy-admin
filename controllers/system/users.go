package system

import (
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/common/global"
	mod "new-go-edas-admin/models/system"
	service "new-go-edas-admin/service/system"
)

// 用户详情
func GetUserInfo(ctx *gin.Context) {
	idStr, _ := ctx.GetQuery("id")
	data, err := service.NewUserInfo().UserInfo(idStr)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 获取登录用户详情
func LoginUserInfo(ctx *gin.Context) {
	id, _ := ctx.Keys["id"]
	data, err := service.NewUserInfo().UserInfo(id.(string))
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)
}

// 用户列表
func UserList(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name"`
		Limit int    `form:"limit"`
		Page  int    `form:"page"`
	})
	if err := ctx.ShouldBind(params); err != nil {
		global.TPLogger.Error("用户查询数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	data, err := service.NewUserInfo().UserList(params.Name, params.Page, params.Limit)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	global.ReturnContext(ctx).Successful("success", data)

}

// 用户更新
func UserUpdate(ctx *gin.Context) {
	params := new(mod.User)
	if err := ctx.ShouldBind(&params); err != nil {
		global.TPLogger.Error("用户更新数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := service.NewUserInfo().UserAdd(params)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "用户添加成功")

}

// 用户添加
func UserAdd(ctx *gin.Context) {
	params := new(mod.User)
	if err := ctx.ShouldBind(&params); err != nil {
		global.TPLogger.Error("用户添加数据绑定失败：", err)
		global.ReturnContext(ctx).Failed("failed", err.Error())
		return
	}
	err := service.NewUserInfo().UserAdd(params)
	if err != nil {
		global.ReturnContext(ctx).Failed("failed", err)
		return
	}
	global.ReturnContext(ctx).Successful("success", "用户添加成功")

}
