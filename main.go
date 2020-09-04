package main

import (
	_ "baseapi/core"
	"baseapi/global"
	"baseapi/routers"
	"fmt"
	"net/http"
)

func main() {
	app := global.BA_CONFIG.App

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", app.HTTPPort),
		Handler:        router,
		ReadTimeout:    app.ReadTimeout,
		WriteTimeout:   app.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
