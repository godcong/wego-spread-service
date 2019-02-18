package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
)

// SpreadList ...
func SpreadList(ver) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model.WechatUserUser(ctx)
	}
}
