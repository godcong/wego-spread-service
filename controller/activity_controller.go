package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
	log "github.com/sirupsen/logrus"
)

// ActivityList 活动列表
func ActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := ctx.Query("type")
		user := model.GetUser(ctx)
		sess := model.Where("is_public = ?", true)
		if t == "user" {
			sess = model.Where("user_id = ?", user.ID)
		}
		act := model.NewActivity("")
		act.IsPublic = true
		var acts []*model.Activity
		e := sess.Find(&acts)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, acts)
	}
}

// ActivityUserList 活动列表
func ActivityUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		act := model.NewActivity("")
		act.UserID = user.ID
		var acts []*model.Activity
		b, e := model.Get(nil, act)
		if e != nil || !b {
			log.Error(b, e)
			Error(ctx, e)
			return
		}
		Success(ctx, acts)
	}
}

// ActivityShow 活动详情
func ActivityShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		activity := model.NewActivity(id)
		_, err := model.Get(nil, activity)
		if err != nil {
			Error(ctx, err)
			return
		}
		Success(ctx, activity)
	}
}
