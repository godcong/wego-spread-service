package service

import "github.com/godcong/wego-auth-manager/config"

// Service ...
type Service struct {
	config *config.Configure
	rest   *HTTPServer
}

var globalService *Service

// New ...
func New(cfg *config.Configure) *Service {
	return &Service{
		config: cfg,
		rest:   NewHTTPServer(cfg),
	}
}

// Start ...
func Start(cfg *config.Configure) {
	globalService = New(cfg)
	globalService.rest.Start()
}

// Stop ...
func Stop() {
	globalService.rest.Stop()
}
