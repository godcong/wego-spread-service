package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-auth-manager/config"

	log "github.com/sirupsen/logrus"
	"net/http"
)

// HTTPServer ...
type HTTPServer struct {
	*gin.Engine
	config *config.Configure
	server *http.Server
	Port   string
}

// NewHTTPServer ...
func NewHTTPServer(cfg *config.Configure) *HTTPServer {
	s := &HTTPServer{
		Engine: gin.Default(),
		config: cfg,
		Port:   config.MustString(cfg.HTTP.Port, ":8080"),
	}
	return s
}

// Start ...
func (s *HTTPServer) Start() {
	s.server = &http.Server{
		Addr:    s.Port,
		Handler: Router(s),
	}
	go func() {
		log.Printf("Listening and serving HTTP on %s\n", s.Port)
		if err := s.server.ListenAndServe(); err != nil {
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
