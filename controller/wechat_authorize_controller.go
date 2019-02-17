package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego"
	"github.com/godcong/wego-spread-service/model"
	"golang.org/x/xerrors"
	"net/http"
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

		account := wego.NewOfficialAccount(p.Config().OfficialAccount)
		url := account.AuthCodeURL(ctx.Query("state"))
		ctx.Redirect(http.StatusFound, url)
	}
}
