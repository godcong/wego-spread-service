package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// UseCrossOrigin ...
func UseCrossOrigin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//for test

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")                          //允许访问所有域
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,token") //header的类型
		ctx.Writer.Header().Set("Content-Type", ctx.Request.Header.Get("Content-Type"))      //返回数据格式是json
		ctx.Writer.WriteHeader(http.StatusOK)
		logrus.Info("cross", ctx.Writer.Header())
	}
}
