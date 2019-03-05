package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-auth-manager/util"

	log "github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"strings"
)

// AuthCheck ...
func AuthCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		var err error
		defer func() {
			if err != nil {
				Error(ctx, err)
				ctx.Abort()
				return
			}
		}()
		if token == "" {
			err = xerrors.New("token is null")
			return
		}
		log.Info("token:", token)
		t, err := util.FromToken(config.Config().WebToken.Key, token)
		if err != nil {
			return
		}
		log.Printf("%+v", t)
		user := model.NewUser(t.UID)

		b, err := user.Get()
		if err != nil {
			return
		}
		if !b {
			err = xerrors.New("no users")
			return
		}

		if !CheckToken(user.Token, token) {
			err = xerrors.New("login expired")
			return
		}
		log.Infof("user:%+v", user)
		ctx.Set("user", user)
		ctx.Next()
	}
}

// CheckToken ...
func CheckToken(src, token string) bool {
	log.Info("compare:", src, token)
	lt := strings.Split(token, ".")
	size := len(lt)
	if size == 0 {
		return false
	}
	if src != lt[size-1] {
		return false
	}
	return true

}
