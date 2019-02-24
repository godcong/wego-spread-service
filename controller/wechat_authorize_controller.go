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
		account.HandleAuthorize(StateHook(code), TokenHook(&code), UserHook(&code, account.AppID, model.WechatTypeH5)).ServeHTTP(ctx.Writer, ctx.Request)
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
func UserHook(code *string, id string, wtype string) wego.UserHook {
	return func(user *core.WechatUserInfo) []byte {
		_, e := model.DB().Transaction(func(session *xorm.Session) (v interface{}, e error) {
			u := (*model.WechatUserInfo)(user)
			if u == nil {
				return nil, xerrors.New("null wechat user")
			}
			defer func() {
				if e != nil {
					log.Error(e)
					e = session.Rollback()
				}
			}()

			var weuser model.WechatUser
			var i int64
			b, e := model.Where("open_id=?", u.OpenID).Get(&weuser)
			model.UserFromHook(&weuser, u, id, wtype)
			if e != nil || !b {
				i, e = model.Insert(nil, &weuser)
			} else {
				i, e = model.Update(nil, weuser.ID, &weuser)
			}
			if e != nil || i == 0 {
				log.Error("wechat user insert or update :", e, i)
				return nil, xerrors.New("wechat user insert or update error")
			}
			var user *model.User
			if b {
				log.Info("user nothing todo")
				return nil, nil
			}
			user = &model.User{
				WechatUserID: weuser.ID,
				UserType:     model.UserTypeUser,
				Nickname:     weuser.Nickname,
				Username:     "user_" + util.GenCRC32(weuser.ID),
				Token:        "",
				Salt:         util.GenerateRandomString(16),
			}
			i, e = model.Insert(session.Clone(), user)
			if e != nil || i == 0 {
				log.Error("user insert:", e, i)
				return nil, xerrors.New("user insert error")
			}

			ua := &model.UserActivity{
				SpreadCode: *code,
			}
			parent, e := ua.CodeSpread(nil)
			if e != nil {
				parent = model.NewSpread("")
			}

			ua.SpreadNumber++
			i, e = model.Update(session.Clone(), ua.ID, ua)
			if e != nil || i == 0 {
				log.Error("user activity update:", e, i)
				return nil, xerrors.New("user activity update error")
			}
			spread := &model.Spread{
				Code:          *code,
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
			i, e = model.Insert(session.Clone(), spread)
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
