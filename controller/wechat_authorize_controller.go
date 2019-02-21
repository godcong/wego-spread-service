package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-spread-service/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

// Authorize ...
func Authorize(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// AuthorizeActivitySpreadNotify ...
func AuthorizeActivitySpreadNotify(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		act := ctx.Param("activity")
		config, e := cache.CachedConfig(act)
		if e != nil {
			log.Error(e)
			Error(ctx, e)
			return
		}
		code := ctx.Query("user") //user spread code
		account := wego.NewOfficialAccount(config.OfficialAccount)
		account.HandleAuthorize(StateHook(code), TokenHook(&code), UserHook(code, account.AppID, 0)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// TokenHook ...
func TokenHook(code *string) wego.TokenHook {
	return func(token *core.Token, state string) []byte {
		*code = cache.GetStateSign(state)
		return nil
	}
}

// UserHook ...
func UserHook(code string, id string, t int) wego.UserHook {
	return func(user *core.WechatUserInfo) []byte {
		_, e := model.DB().Transaction(func(session *xorm.Session) (v interface{}, e error) {
			weuser := model.UserFromHook(user, id, t)
			defer func() {
				if e != nil {
					log.Error(e)
					e = session.Rollback()
				}
			}()
			if weuser == nil {
				return nil, xerrors.New("null wechat user")
			}

			user := &model.User{
				Model:    model.Model{},
				UserType: model.UserTypeUser,
				Nickname: weuser.Nickname,
				//Sign:     util.CRC32(weuser.ID),
				Token: "",
				Salt:  util.GenerateRandomString(16),
			}
			i, e := model.Insert(nil, user)
			if e != nil || i == 0 {
				log.Error("user insert:", e, i)
				return nil, xerrors.New("user insert error")
			}

			ua := &model.UserActivity{
				SpreadCode: code,
			}
			parent, e := ua.CodeSpread()
			if e != nil {
				return nil, e
			}
			weuser.UserID = user.ID
			i, e = model.Insert(nil, weuser)
			if e != nil || i == 0 {
				log.Error("wechat user insert:", e, i)
				return nil, xerrors.New("wechat user insert error")
			}

			ua.SpreadNumber++
			i, e = model.Update(nil, ua.ID, ua)
			if e != nil || i == 0 {
				log.Error("user activity update:", e, i)
				return nil, xerrors.New("user activity update error")
			}
			spread := &model.Spread{
				Code:          code,
				UserID:        user.ID,
				ParentUserID1: parent.UserID,
				ParentUserID2: parent.ParentUserID1,
				ParentUserID3: parent.ParentUserID2,
				ParentUserID4: parent.ParentUserID3,
				ParentUserID5: parent.ParentUserID4,
				ParentUserID6: parent.ParentUserID5,
				ParentUserID7: parent.ParentUserID6,
				ParentUserID8: parent.ParentUserID7,
				ParentUserID9: parent.ParentUserID8,
			}
			i, e = model.Insert(nil, spread)
			if e != nil || i == 0 {
				log.Error("spread insert", i)
				return nil, xerrors.New("spread insert error")
			}
			return nil, nil
		})
		if e != nil {
			log.Error(e)
		}
		return nil
	}
}

// StateHook ...
func StateHook(code string) wego.StateHook {
	return func() string {
		key := util.GenMD5(uuid.New().String())
		cache.SetStateSign(key, code)
		return key
	}
}
