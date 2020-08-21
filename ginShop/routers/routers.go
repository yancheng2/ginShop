package routers

import (
	"ginShop/pkg/setting"
	v1 "ginShop/routers/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine{
	engine := gin.New()
	engine.Use(gin.Logger())			//gin框架中  打印中间件
	engine.Use(gin.Recovery())			//gin框架中恢复中间件，可以从任何恐慌中回复 写入500

	gin.SetMode(setting.ServerSetting.RunMode)//设置运行模式，debug或release

	group := engine.Group("/api/v1")
	{
		group.GET("test",v1.GetAppVersionTest)
	}
	return engine

}
