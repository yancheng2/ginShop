package main

import (
	"fmt"
	"ginShop/models"
	"ginShop/pkg/logging"
	"ginShop/pkg/setting"
	"ginShop/routers"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){
	log.Println("api server start ...")

	setting.SetUp()			//初始化app.ini配置文件
	logging.SetUp()			//初始化日志文件
	models.SetUp()			//初始化数据库

	initRouter := routers.InitRouter()
	//http参数设置
	server := &http.Server{
		Addr:fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),					//设置端口号
		Handler: initRouter,													//gin实例化好的路由组
		ReadTimeout: setting.ServerSetting.ReadTimeout,							//允许读取的最大时间
		WriteTimeout: setting.ServerSetting.WriteTimeout,						//允许写入的最大时间
		MaxHeaderBytes: 1<<20,
	}


	go func() {
		if err := server.ListenAndServe();err != nil{
			log.Printf("Listen:%s\n",err)
		}
	}()

	// 设置发送信号通知的通道
	// 我们必须使用缓冲通道，否则可能会丢失信号
	// 如果我们还没准备好接收信号
	//os.signal包实现对信号的处理  notify方法用来监听收到的信号   stop方法用来取消监听
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	<-quit

	log.Println("shutdown server ...")

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()  //取消协程操作

	if err := server.Shutdown(timeout);err != nil{
		log.Fatal("server shutdown:",err)
	}
	log.Println("程序服务关闭退出")
}
