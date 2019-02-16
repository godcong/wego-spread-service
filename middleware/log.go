package middleware

import (
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

// VisitLog ...
func VisitLog(ver string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

	}
}

// RemoteIP ...
func RemoteIP(ctx *fasthttp.RequestCtx) {
	host := string(ctx.Request.Header.Peek("REMOTE-HOST"))
	if host == "" {
		if ip := ctx.RemoteIP(); ip != nil {
			host = ip.String()
		}
	}
	log.Info("host:", host)
}
