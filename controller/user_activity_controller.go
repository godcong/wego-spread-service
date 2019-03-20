package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego-auth-manager/model"
)

// StringToBool ...
func StringToBool(s string) bool {
	return s == "true"
}

// UserActivityList 我参加的活动
func UserActivityList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.GetUser(ctx)
		favorite := ctx.Query("favorite")
		act := model.NewUserActivity("")
		act.UserID = user.ID

		var session *xorm.Session
		if StringToBool(favorite) {
			session = model.Where("user_activity.is_favorite = ?", true)
			//} else if favorite == "false" {
			//session = model.Where("user_activity.is_favorite = ?", false)
			//} else {
			//favorite == ""
		}
		activities, e := act.Activities(session)
		if e != nil {
			Error(ctx, e)
			return
		}
		Success(ctx, activities)
	}
}

// UserActivityFavoriteUpdate ...
func
UserActivityFavoriteUpdate(ver
string) gin.HandlerFunc{
return func (ctx *gin.Context){
user := model.GetUser(ctx)
id := ctx.Param("id")
favorite := ctx.PostForm("favorite")
act := model.NewUserActivity(id)
act.UserID = user.ID
b, e := model.Get(nil, act)
if e != nil || !b{
log.Error(e, b)
Error(ctx, xerrors.New("activity not found!"))
return
}
act.IsFavorite = false
if favorite == "true"{
act.IsFavorite = true
} else if favorite == "false"{
act.IsFavorite = false
}

//i, e := act.Update("is_favorite")
i, e := model.UpdateWithColumn(nil, act.ID, act, "is_favorite")
if e != nil || i == 0{
log.Error(e, i)
Error(ctx, xerrors.New("activity update error"))
return
}
Success(ctx, act)
}
}

// UserActivityCodeGet 活动分享
func
UserActivityCodeGet(ver
string) gin.HandlerFunc{
return func (ctx *gin.Context){
id := ctx.Param("id")
user := model.GetUser(ctx)
act := model.NewActivity(id)
act.UserID = user.ID

p, e := act.Property()
if e != nil{
Error(ctx, e)
return
}
jssdk := wego.NewJSSDK(p.Config().JSSDK)
config := jssdk.BuildConfig("")
if config == nil{
Error(ctx, xerrors.New("null config result"))
return
}
Success(ctx, config)
}
}

// UserActivityJoin 我申请的活动
func
UserActivityJoin(ver
string) gin.HandlerFunc{
return func (ctx *gin.Context){
user := model.GetUser(ctx)
code := ctx.Param("code")
act := model.Activity{
Code: code,
}
b, e := act.Get()
if e != nil || !b{
log.Error(e, b)
Error(ctx, xerrors.New("activity not found"))
return
}
property, e := act.Property()
if e != nil{
Error(ctx, xerrors.New("property not found"))
return
}
verified := false
if !act.NeedVerify{
verified = true
}
ua := model.UserActivity{
ActivityID: act.ID,
UserID:     user.ID,
IsVerified: verified,
PropertyID: property.ID,
SpreadCode: util.GenCRC32(act.ID + user.ID),
}
i, e := model.Insert(nil, &ua)
if e != nil || i == 0{
log.Error(e, b)
Error(ctx, xerrors.New("user activity insert error"))
return
}
Success(ctx, ua)
}
}
