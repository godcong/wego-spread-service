//go:generate statik -f -src=./dist
package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/go-trait"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-spread-service/cache"
	"github.com/godcong/wego-spread-service/service"
	_ "github.com/godcong/wego-spread-service/statik"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "load config from path")
var logPath = flag.String("log", "logs/spread.log", "set log name")
var elk = flag.Bool("elk", false, "set to open the elk")

func main() {
	flag.Parse()

	if *elk {
		trait.InitElasticLog("wego-spread-service")
	} else {
		trait.InitRotateLog(*logPath)
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.InitConfig(*configPath)
	log.Infof("%+v", cfg)
	model.InitDB(cfg)

	c := cache.DefaultCache()
	cache.InitWegoCache(c)
	cache.InitStateCache(c)
	cache.InitPropertyCache(c)

	//start
	service.Start(cfg)
	//start
	go func() {
		sig := <-sigs
		//bm.Stop()
		fmt.Println(sig, "exiting")
		done <- true
	}()
	<-done
}

func initLog() {
	//client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	//if err != nil {
	//	log.Panic(err)
	//}
	//
	//t, err := elogrus.NewElasticHook(client, "localhost", log.TraceLevel, "ipfs-cluster-monitor")
	//if err != nil {
	//	log.Panic(err)
	//}
	//log.AddHook(t)
}
