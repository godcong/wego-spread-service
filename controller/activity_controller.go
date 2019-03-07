package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/model"
)

// ActivityList 活动列表
func ActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		act := model.NewActivity("")
		act.IsPublic = true
		var acts []*model.Activity
		e := model.Where("is_public = ?", true).Find(&acts)
		if e != nil {
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
