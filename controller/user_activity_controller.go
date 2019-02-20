package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

// UserActivityList 我的活动
func UserActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO
		user := model.GetUser(ctx)
		log.Error(user)
	}
}

// UserActivityJoin 活动申请
func UserActivityJoin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//TODO
		user := model.GetUser(ctx)
		log.Error(user)
	}
}
