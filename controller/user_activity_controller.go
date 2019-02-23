package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego/util"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// UserActivityList 我的活动
func UserActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		page := model.PageUserActivity(model.ParsePaginate(ctx.Request.URL.Query()))
		e := page.PageWhere(model.Where("user_id = ?", user.ID))

		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, page)
	}
}

// UserActivityJoin 活动申请
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
		ua := model.UserActivity{
			ActivityID: act.ID,
			UserID:     user.ID,
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
