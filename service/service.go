package service

import "github.com/godcong/wego-auth-manager/config"

// Service ...
type Service struct {
	config *config.Configure
	http   *HTTPServer
}

var global *Service

// New ...
func New(cfg *config.Configure) *Service {
	return &Service{
		config: cfg,
		http:   NewHTTPServer(cfg),
	}
}

// Start ...
func Start(cfg *config.Configure) {
	global = New(cfg)
	global.http.Start()
}

// Stop ...
func Stop() {
	global.http.Stop()
}
