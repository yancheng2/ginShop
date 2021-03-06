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
			tokenG.POST("addGoodsToCart", v1.AddGoodsToCart)    //加购
			tokenG.GET("getCartGoodsList", v1.GetCartGoodsList) //购物车列表
			tokenG.GET("delCartGoods", v1.DelCartGoods)         //购物车列表
		}

		//无需token登录的接口
		group.POST("login", v1.Login)              //登录
		group.GET("goodsList", v1.GetGoodsList)    //商品列表
		group.GET("goodsDetails", v1.GoodsDetails) //商品详情
		group.GET("categoryList", v1.GetCategory)  //分类树

		group.GET("SyncGoodsToEs", v1.SyncGoodsToEs) //同步es
		group.GET("CreateIndex", v1.CreateIndex)     //创建index
	}
	return engine

}
