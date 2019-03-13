package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-spread-service/controller"
	"github.com/godcong/wego-spread-service/middleware"
)

// Router ...
func Router(server *HTTPServer) *gin.Engine {
	version := "v0"
	eng := server.Engine
	eng.Use(middleware.UseCrossOrigin(version))

	v0 := eng.Group("api").Group(version)
	v0.GET("authorize/:activity/*uri", controller.AuthorizeActivitySpreadNotify(version))

	spreadN := v0.Group("spread")
	spreadN.GET("activity/:id", controller.ActivityShow(version))

	spreadA := v0.Group("spread", middleware.AuthCheck(version))

	spreadA.GET("activity", controller.ActivityList(version))
	spreadA.GET("user/info", controller.UserInfo(version))
	spreadA.GET("user/activity/:param", controller.UserActivityList(version))
	spreadA.GET("user/spread", controller.UserSpreadList(version))
	spreadA.POST("user/activity/:code", controller.UserActivityJoin(version))
	spreadA.GET("activity/:id/share", controller.UserActivityShareGet(version))
	spreadA.GET("spread/:id/share", controller.UserSpreadShareGet(version))
	return eng
}

func isInstalled() bool {
	return false
}

// AccessControlAllow ...
func AccessControlAllow(ctx *gin.Context) {
	origin := ctx.Request.Header.Get("origin")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, XMLHttpRequest, "+
		"Accept-Encoding, X-CSRF-Token, Authorization")
	if ctx.Request.Method == "OPTIONS" {
		ctx.String(200, "ok")
		return
	}
	ctx.Next()
}
