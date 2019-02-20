package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

// ActivityShare 活动分享
func ActivityShare(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO
		user := model.GetUser(ctx)
		log.Error(user)
	}
}
