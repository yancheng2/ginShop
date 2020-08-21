package routers

import (
	"ginShop/middleware"
	"ginShop/pkg/setting"
	v1 "ginShop/server/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())   //gin框架中  全局中间件-日志
	engine.Use(gin.Recovery()) //gin框架中恢复中间件，可以从任何恐慌中回复 写入500

	gin.SetMode(setting.ServerSetting.RunMode) //设置运行模式，debug或release

	group := engine.Group("/api/v1")
	{
		tokenG := group.Group("") //再次分组
		//需要token登录的接口
		tokenG.Use(middleware.TokenVer())
		{

		}

		//无需token登录的接口
		group.POST("login", v1.Login)
	}
	return engine

}
