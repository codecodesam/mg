package main

import (
	"net/http"
	"os"
	"time"

	"github.com/codecodesam/mg/micro/auth/api/router"
	"github.com/codecodesam/mg/pkg/config"
	"github.com/codecodesam/mg/pkg/logger"
	mw "github.com/codecodesam/mg/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// init logger
	initLogger()
	// init config
	initConfig()
	// init web gin
	initWeb()
	// close logger and flush content to file system
	closeLogger()
}

func initWeb() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(mw.LoggerMiddleware(), mw.RecoveryMiddleware(true))
	router.Register(engine)

	// load config
	addr, getAddrErr := config.GetStringValue(config.CFG_IP_PORT)
	if getAddrErr != nil {
		panic(getAddrErr)
	}
	rt := config.GetIntValueWithDefaultValue(config.CFG_SERVER_READ_TIMEOUT, 60)
	wt := config.GetIntValueWithDefaultValue(config.CFG_SERVER_WRITE_TIMEOUT, 60)

	server := &http.Server{
		Addr:           addr,
		Handler:        engine,
		ReadTimeout:    time.Duration(rt) * time.Second,
		WriteTimeout:   time.Duration(wt) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func initConfig() {
	config.NewConfigManager()
}

func initLogger() {
	path := os.Getenv(config.ENV_LOG_PATH)
	if path == "" {
		panic("can't find logger path")
	}
	logger.InitLogger(path)
}

func closeLogger() {
	err := logger.Log.Sync()
	if err != nil {
		// ignore this error
	}
}
