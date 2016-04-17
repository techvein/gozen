package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"

	"gozen/config"
	"gozen/controllers"
)

func main() {
	var (
		pprof   bool
		useHttp bool
	)

	flag.BoolVar(&pprof, "pprof", false, "プロファイリング結果をポート6060で確認できるようになります。")
	flag.BoolVar(&useHttp, "http", false, "Goのhttpサーバーを使用します。")

	flag.Parse()

	lf, err := os.OpenFile(config.Log.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		// logger.Init前のためlogを使用
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	logger.Init("gozenLogger", config.Log.Verbose, true, lf)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	switch config.GetEnv() {
	case config.Production:
		// TODO
		gin.SetMode(gin.ReleaseMode)
		logger.Infoln(config.ProductionStr)
	case config.Staging:
		// TODO
		logger.Infoln(config.StagingStr)
	case config.Development:
		// TODO
		logger.Infoln(config.DevelopmentStr)
	}

	if pprof {
		runtime.SetBlockProfileRate(1)
		go func() {
			// :6060/debug/pprof/ でプロファイリング結果を確認できる
			logger.Infoln(http.ListenAndServe("0.0.0.0:6060", nil))
		}()
	}

	router := controllers.Routes()

	if useHttp {
		logger.Infoln("use http")
		go func() {
			http.ListenAndServe(":9000", router)
		}()
	} else {
		logger.Infoln("use fcgi(nginx)")
		listen, err := net.Listen("tcp", "127.0.0.1:9000")
		if err != nil {
			return
		}
		go func() {
			fcgi.Serve(listen, router)
		}()
	}

	<-sig
}
