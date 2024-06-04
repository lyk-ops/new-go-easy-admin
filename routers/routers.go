package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"new-go-edas-admin/common/global"
	"new-go-edas-admin/middles"
	"new-go-edas-admin/routers/system"
	"time"
)

func BaseRouters() *gin.Engine {
	r := gin.New()
	// 自定义日志格式
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s | method: %s | path: %s | host: %s | proto: %s | code: %d | %s | %s ]\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
		)
	}))
	// 开启全部跨域允许
	r.Use(middles.Cors())
	authMiddleware, err := middles.InitAuth()
	if err != nil {
		global.TPLogger.Error("初始化JWT中间件失败: ", err)
		panic(fmt.Sprintf("初始化JWT中间件失败：%v", err))
	}
	// 健康检查
	r.GET("/health", func(ctx *gin.Context) {
		global.ReturnContext(ctx).Successful("success", "success")
		return
	})
	// 不需要做鉴权的接口 PublicGroup
	PublicGroup := r.Group("/api/base")
	{
		system.InitBaseRouters(PublicGroup, authMiddleware)
	}
	// 需要做鉴权的接口
	PrivateGroup := r.Group("/api/system")
	{
		system.InitUserRouters(PrivateGroup, authMiddleware)
	}
	return r
}
