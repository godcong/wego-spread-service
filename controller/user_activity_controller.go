package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserActivityList 我参加的活动
func UserActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		cas := ctx.Param("favorite")
		act := model.NewUserActivity("")
		act.UserID = user.ID

		var session *xorm.Session
		if cas == "favorite" {
			session = model.Where("user_activity.is_favorite = ?", true)
		}
		activities, e := act.Activities(session)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, activities)
	}
}

// UserActivityFavorite ...
func UserActivityFavorite(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UserActivityShareGet 活动分享
func UserActivityShareGet(ver string) gin.HandlerFunc {
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

// UserActivityJoin 我申请的活动
func UserActivityJoin(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		code := ctx.Param("code")
		act := model.Activity{
			Code: code,
		}
		b, e := act.Get()
		if e != nil || !b {
			log.Error(e, b)
			Error(ctx, xerrors.New("activity not found"))
			return
		}
		verified := false
		if !act.NeedVerify {
			verified = true
		}
		ua := model.UserActivity{
			ActivityID: act.ID,
			UserID:     user.ID,
			IsVerified: verified,
			SpreadCode: util.GenCRC32(act.ID + user.ID),
		}
		i, e := model.Insert(nil, &ua)
		if e != nil || i == 0 {
			log.Error(e, b)
			Error(ctx, xerrors.New("user activity insert error"))
			return
		}
		Success(ctx, ua)
	}
}
