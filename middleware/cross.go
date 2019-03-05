package middleware

import "github.com/gin-gonic/gin"

// UseCrossOrigin ...
func UseCrossOrigin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//for test
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")                   //允许访问所有域
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,token") //header的类型
		ctx.Writer.Header().Set("Content-Type", "application/json")                   //返回数据格式是json
	}
}
