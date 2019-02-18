package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego-spread-service/cache"
	"github.com/godcong/wego-spread-service/model"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/util"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

// Authorize ...
func Authorize(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// AuthorizeNotify ...
func AuthorizeNotify(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sign := ctx.Param("sign")
		config := cache.GetSignConfig(sign)
		if config == nil {
			p := model.Property{
				Sign: sign,
			}
			b, e := model.Get(nil, &p)
			if e != nil {
				Error(ctx, e)
				return
			}
			if !b {
				Error(ctx, xerrors.New("no found"))
				return
			}
			config = p.Config()
			cache.SetSignConfig(sign, config)
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
		wu := model.UserFromHook(user, id, t)
		if wu == nil {
			log.Error(xerrors.New("null wechat user"))
			return nil
		}
		i, e := model.Insert(nil, wu)
		if e != nil || i == 0 {
			log.Error(e, i)
			return nil
		}
		parent := &model.Spread{
			SelfSign: userSign,
		}
		b, e := model.Get(nil, parent)
		if e != nil || !b {
			log.Error(e, b)
			//continue
		}
		spread := &model.Spread{
			WechatUserID: wu.ID,
			SelfSign:     wu.Sign,
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
			log.Error(e, i)
			return nil
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
