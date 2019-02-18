package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego-spread-service/cache"
	"github.com/godcong/wego-spread-service/model"
	"github.com/godcong/wego/core"
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
		account.HandleAuthorize(StateHook(userSign), TokenHook(&userSign), UserHook(userSign)).ServeHTTP(ctx.Writer, ctx.Request)
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
		model.UserFromHook(user, id, t)
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
