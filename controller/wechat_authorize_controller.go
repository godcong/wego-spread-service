package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego"
	"github.com/godcong/wego-spread-service/cache"
	"github.com/godcong/wego-spread-service/model"
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

// AuthorizeSignNotify ...
func AuthorizeSignNotify(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sign := ctx.Param("sign")
		config, e := model.CachedConfig(sign)
		if e != nil {
			Error(ctx, e)
		}
		userSign := ctx.Query("user")
		account := wego.NewOfficialAccount(config.OfficialAccount)
		account.HandleAuthorize(StateHook(userSign), TokenHook(&userSign), UserHook(userSign, account.AppID, 0)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// TokenHook ...
func TokenHook(userSign *string) wego.TokenHook {
	return func(token *core.Token, state string) []byte {
		*userSign = cache.GetStateSign(state)
		return nil
	}
}

// UserHook ...
func UserHook(userSign string, id string, t int) wego.UserHook {
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
			i, e := model.Insert(nil, weuser)
			if e != nil || i == 0 {
				log.Error("wechat user insert:", e, i)
				return nil, xerrors.New("wechat user insert error")
			}

			user := &model.User{
				Model:        model.Model{},
				WechatUserID: weuser.ID,
				UserType:     model.UserTypeUser,
				Nickname:     weuser.Nickname,
				Sign:         util.CRC32(weuser.ID),
				Token:        "",
				Salt:         util.GenerateRandomString(16),
			}
			i, e = model.Insert(nil, user)
			if e != nil || i == 0 {
				log.Error("user insert:", e, i)
				return nil, xerrors.New("user insert error")
			}

			parent := &model.Spread{
				SelfSign: userSign,
			}
			b, e := model.Get(nil, parent)
			if e != nil || !b {
				log.Error("parent:", b)
				//continue
			}
			spread := &model.Spread{
				WechatUserID: weuser.ID,
				SelfSign:     user.Sign,
				ParentSign:   userSign,
				ParentSign2:  parent.ParentSign,
				ParentSign3:  parent.ParentSign2,
				ParentSign4:  parent.ParentSign3,
				ParentSign5:  parent.ParentSign4,
				ParentSign6:  parent.ParentSign5,
				ParentSign7:  parent.ParentSign6,
				ParentSign8:  parent.ParentSign7,
				ParentSign9:  parent.ParentSign8,
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
func StateHook(userSign string) wego.StateHook {
	return func() string {
		key := util.GenMD5(uuid.New().String())
		cache.SetStateSign(key, userSign)
		return key
	}
}
