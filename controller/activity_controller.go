package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

// ActivityShare 活动分享
func ActivityShare(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.GetUser(ctx)
		act := model.NewActivity(id)
		act.UserID = user.ID
		b, e := act.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, e)
			return
		}
		act.CodeProperty()
		Success(ctx, act.Code)
	}
}
