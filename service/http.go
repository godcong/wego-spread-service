package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net/http"
)

// HTTPServer ...
type HTTPServer struct {
	*gin.Engine
	loader *RouteLoader
	config *config.Configure
	server *http.Server
	Port   string
}

// NewHTTPServer ...
func NewHTTPServer(cfg *config.Configure) *HTTPServer {
	s := &HTTPServer{
		config: cfg,
		Port:   config.MustString(cfg.REST.Port, ":8080"),
		loader: NewRouteLoader("v0"),
	}
	return s
}

// Start ...
func (s *HTTPServer) Start() {
	go func() {
		log.Printf("Listening and serving HTTP on %s\n", s.Port)
		if err := fasthttp.ListenAndServe(s.Port, s.loader.router); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

}

// Stop ...
func (s *HTTPServer) Stop() {
	if err := s.server.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
