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
