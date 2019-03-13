package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	"golang.org/x/xerrors"
)

// UserSpreadList 我的推广
func UserSpreadList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		spread := model.NewSpread("")
		activities, e := spread.SpreadActivity(model.Where("spread.parent_user_id_1 = ?", user.ID))
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, activities)
	}
}

// UserSpreadShareGet 我的分享码
func UserSpreadShareGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		user := model.GetUser(ctx)
		act := model.NewUserActivity(id)
		act.UserID = user.ID
		p, e := act.Property(model.Where("user_activity.verified = ?", true))
		if e != nil {
			Error(ctx, e)
			return
		}
		url, b := ctx.GetQuery("url")
		if !b {
			url = p.Host
		}
		jssdk := wego.NewJSSDK(p.Config().JSSDK, wego.JSSDKOption{
			URL: url,
		})
		config := jssdk.BuildConfig("")
		if config == nil {
			Error(ctx, xerrors.New("null config result"))
			return
		}
		Success(ctx, config)
	}
}
