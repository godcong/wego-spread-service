package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

// Router ...
func Router(eng *gin.Engine) {

	verV0 := "v0"
	eng.NoRoute(func(ctx *gin.Context) {
		log.Println("no route")
		dir, file := path.Split(ctx.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			ctx.File("./dist/index.html")
		} else {
			ctx.File("./dist/" + path.Join(dir, file))
		}
	})

	//eng.Static("webui", "dist")
	eng.Use(AccessControlAllow)
	g0 := eng.Group(verV0)

	g0.GET("inited", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Code": 0})
	})

	////登录
	//g0.POST("login", LoginPOST(verV0))
	////组织注册
	//g0.POST("register", RegisterPOST(verV0))
	//
	//g0.POST("genesis", GenesisGet(verV0))

}

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
