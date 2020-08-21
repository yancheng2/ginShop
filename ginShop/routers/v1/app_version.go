package v1

import (
	"ginShop/pkg/e"
	"github.com/gin-gonic/gin"
)

func GetAppVersionTest(c *gin.Context){
	c.JSON(200,gin.H{
		"Code":200,
		"Msg":e.GetMsg(200),
		"data":"测试成功",
	})
}
