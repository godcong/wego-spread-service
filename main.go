//go:generate statik -f -src=./dist
package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/wego-auth-manager/config"
	"github.com/godcong/wego-auth-manager/log"
	"github.com/godcong/wego-auth-manager/model"
	"github.com/godcong/wego-spread-service/service"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "load config from path")
var logPath = flag.String("log", "spread.log", "set log name")

func main() {
	flag.Parse()

	log.InitLog("wego-spread-service")
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	cfg := config.InitConfig(*configPath)
	model.InitDB(cfg)

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
