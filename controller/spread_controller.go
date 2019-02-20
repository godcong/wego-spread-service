package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"

	log "github.com/sirupsen/logrus"
)

// SpreadList 推广列表
func SpreadList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		log.Error(user)
	}
}
