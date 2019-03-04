package middleware

import "github.com/gin-gonic/gin"

// UseCrossOrigin ...
func UseCrossOrigin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//for test
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		//w.Header().Set("content-type", "application/json")             //返回数据格式是json
	}
}
