package service

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/godcong/wego-manager-service/config"
	"log"
	"net/http"
)

// service ...
type service struct {
	config *config.Configure
	eng    *gin.Engine
	server *http.Server
}

var global *service
var configPath = flag.String("config", "config.toml", "load config from path")

func init() {
	flag.Parse()
	cfg := config.InitLoader(*configPath)
	global = initService(cfg)
}

// Start ...
func Start() {
	Router(global.eng)

	go func() {
		log.Printf("[GIN-debug] Listening and serving HTTP on %s\n", global.server.Addr)
		if err := global.server.ListenAndServe(); err != nil {
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

}

// Stop ...
func Stop() error {
	return global.server.Shutdown(nil)
}

func initService(cfg *config.Configure) *service {
	eng := gin.Default()
	return &service{
		eng: eng,
		server: &http.Server{
			Addr:    ":7788",
			Handler: eng,
		},
	}
}
