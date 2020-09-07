package main

import (
	_ "baseapi/core"
	"baseapi/global"
	"baseapi/routers"
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"
)

func main() {
	app := global.BA_CONFIG.App

	endless.DefaultReadTimeOut = app.ReadTimeout
	endless.DefaultWriteTimeOut = app.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", app.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
