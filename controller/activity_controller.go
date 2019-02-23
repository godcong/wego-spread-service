package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

// ActivityShareGet 活动分享
func ActivityShareGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.GetUser(ctx)
		act := model.NewActivity(id)
		act.UserID = user.ID

		p, e := act.Property()
		if e != nil {
			log.Error(e)
			Error(ctx, e)
			return
		}
		jssdk := wego.NewJSSDK(p.Config().JSSDK)
		config := jssdk.BuildConfig("")
		Success(ctx, config)
	}
}
