//go:generate statik -f -src=./dist
package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/godcong/wego-spread-service/config"
	"github.com/godcong/wego-spread-service/model"
	"github.com/godcong/wego-spread-service/service"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/signal"
	"syscall"
)

var configPath = flag.String("config", "config.toml", "load config from path")
var logPath = flag.String("log", "spread.log", "set log name")

func main() {
	flag.Parse()
	file, err := os.OpenFile(*logPath, os.O_SYNC|os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFormatter(&log.JSONFormatter{})
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
