package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	"github.com/godcong/wego"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	ut "github.com/godcong/wego-auth-manager/util"
	"github.com/godcong/wego-spread-service/cache"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/util"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"net/http"
	"net/url"
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
		account.HandleAuthorize(StateHook(ctx, code), TokenHook(ctx, &code), UserHook(ctx, &code, account.AppID, model.WechatTypeH5)).ServeHTTP(ctx.Writer, ctx.Request)
	}
}

// TokenHook ...
func TokenHook(ctx *gin.Context, code *string) wego.TokenHook {
	return func(w http.ResponseWriter, req *http.Request, token *core.Token, state string) []byte {
		*code = cache.GetStateSign(state)
		return nil
	}
}

// UserHook ...
func UserHook(ctx *gin.Context, code *string, id string, wtype string) wego.UserHook {
	return func(w http.ResponseWriter, req *http.Request, user *core.WechatUserInfo) []byte {
		token, e := model.DB().Transaction(func(session *xorm.Session) (v interface{}, e error) {
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
				user = &model.User{
					WechatUserID: weuser.ID,
				}
				b, e = user.Get()
				if e != nil || !b {
					log.Error(e, b)
					return nil, xerrors.New("user not found")
				}
				token, e := user.Login()
				if e != nil {
					return nil, e
				}
				i, e = user.Update("token")
				if e != nil || i == 0 {
					log.Error(e, i)
					return nil, xerrors.New("login error")
				}
				return ut.ToToken(config.Config().WebToken.Key, token)
			}
			user = &model.User{
				WechatUserID: weuser.ID,
				UserType:     model.UserTypeUser,
				Nickname:     weuser.Nickname,
				Username:     "USER_" + util.GenCRC32(weuser.ID),
				Salt:         util.GenerateRandomString(16),
			}
			//do login before insert
			token, e := user.Login()
			if e != nil {
				return nil, e
			}
			i, e = model.Insert(session, user)
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
			i, e = model.Update(session, ua.ID, ua)
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

			i, e = model.Insert(session, spread)
			if e != nil || i == 0 {
				log.Error("spread insert", i)
				return nil, xerrors.New("spread insert error")
			}
			return ut.ToToken(config.Config().WebToken.Key, token)
		})
		if e != nil {
			log.Error(e)
		}

		if v, b := token.(string); b {
			log.Info("redirect")
			v := url.Values{
				"token": {v},
			}
			http.Redirect(w, req, "http://localhost/index.html?"+v.Encode(), http.StatusFound)
		}

		return nil
	}
}

// StateHook ...
func StateHook(ctx *gin.Context, code string) wego.StateHook {
	return func(w http.ResponseWriter, req *http.Request) string {
		key := util.GenMD5(uuid.New().String())
		cache.SetStateSign(key, code)
		return key
	}
}
