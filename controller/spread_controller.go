package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-spread-service/model"
	log "github.com/sirupsen/logrus"
)

// SpreadList ...
func SpreadList(ver) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		weuser := model.GetWechatUser(ctx)
		log.Error(weuser)
	}
}
