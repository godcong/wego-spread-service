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

	spreads := v0.Group("spread", middleware.AuthCheck(version))
	spreads.GET("activities", controller.ActivityList(version))
	spreads.GET("activities/:id", controller.ActivityShow(version))
	spreads.GET("activities/:id/share", controller.UserActivityShareGet(version))
	spreads.GET("users/info", controller.UserInfo(version))
	spreads.GET("users/activity/show/:favorite", controller.UserActivityList(version))
	spreads.GET("users/spread", controller.UserSpreadList(version))
	spreads.POST("users/activity/:id/favorite/:status", controller.UserActivityFavorite(version))
	spreads.POST("users/activity/:id/join/:code", controller.UserActivityJoin(version))
	spreads.GET("spreads/:id/share", controller.UserSpreadShareGet(version))
	return eng
}
