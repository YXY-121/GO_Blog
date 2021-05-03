package main

import (
	"awesomeProject/global"
	"awesomeProject/internal/model"
	"awesomeProject/internal/routers"
	setting "awesomeProject/pkg/setting"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)
func init(){
	err:=setupSetting()
	if err!=nil{
		log.Fatal("init.setupSetting err:%v",err)
	}
}

func setupSetting() error {
	setting,err :=setting.NewSetting()
	if err!=nil{
		return nil
	}
	err=setting.ReadSection("Server",&global.ServerSetting)
	if err!=nil{
		return nil
	}


	err=setting.ReadSection("App",&global.AppSetting)
	if err!=nil{
		return nil
	}

	err=setting.ReadSection("Database",&global.DatabaseSetting)
	if err!=nil{
		return nil
	}
	global.ServerSetting.ReadTimeout *=time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil

}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-bookcd
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router :=routers.NewRouter()
	s:=&http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

	}

func setupDBEngine() error {
	var err error
	global.DBEngine,err=model.NewDBEngine(global.DatabaseSetting)
	if(err!=nil){
		return err
	}
	return nil
}
