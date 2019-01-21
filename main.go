package main

import (
	"context"
	"fmt"
	"gin-example/config"
	"gin-example/models"
	"gin-example/mylog"
	"gin-example/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	config.Setup()
	mylog.SetUp()
	models.Setup()

	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    config.ServerSetting.ReadTimeout,
		WriteTimeout:   config.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	//优雅关闭，热更新根据场景设计
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
