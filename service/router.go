package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-spread-service/controller"
	"github.com/godcong/wego-spread-service/middleware"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

// Router ...
func Router(server *HTTPServer) *gin.Engine {
	version := "v0"
	eng := server.Engine
	//eng.Use(middleware.UseCrossOrigin(version))
	st, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	eng.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/webui")
	})
	eng.StaticFS("webui", st)
	eng.StaticFile("MP_verify_Z0o2ocg9NBJ4cDFG.txt", "./static/MP_verify_Z0o2ocg9NBJ4cDFG.txt")
	v0 := eng.Group("api").Group(version)
	v0.GET("authorize/:activity/*uri", controller.AuthorizeActivitySpreadNotify(version))

	spreadN := v0.Group("spread")
	spreadN.GET("activity/:id", controller.ActivityShow(version))

	spreadA := v0.Group("spread", middleware.AuthCheck(version))

	spreadA.GET("activity", controller.ActivityList(version))
	spreadA.GET("user/info", controller.UserInfo(version))
	spreadA.GET("user/activity/show/:favorite", controller.UserActivityList(version))
	spreadA.POST("user/activity/:id/favorite/:status", controller.UserActivityFavorite(version))
	spreadA.GET("user/spread", controller.UserSpreadList(version))
	spreadA.POST("user/activity/:id/join/:code", controller.UserActivityJoin(version))
	spreadA.GET("activity/:id/share", controller.UserActivityShareGet(version))
	spreadA.GET("spread/:id/share", controller.UserSpreadShareGet(version))
	return eng
}
