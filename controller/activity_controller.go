package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	"golang.org/x/xerrors"
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
			Error(ctx, e)
			return
		}
		jssdk := wego.NewJSSDK(p.Config().JSSDK)
		config := jssdk.BuildConfig("")
		if config == nil {
			Error(ctx, xerrors.New("null config result"))
			return
		}
		Success(ctx, config)
	}
}
