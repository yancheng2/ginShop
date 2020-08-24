package v1

import (
	"strconv"
	"ginShop/models"
	"ginShop/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetGoodsList(c *gin.Context){
	page_s := c.Query("page")
	cat_id := c.Query("cat_id")
	is_hot := c.Query("is_hot")
	is_sale := c.Query("is_sale")

	if page_s==""{
		page_s= "1"
	}
	page,_ := strconv.Atoi(page_s)
	offset := (page-1)*20

	var where map[string]interface{}
	where = make(map[string]interface{})
	where["page"]=page
	where["cat_id"]=cat_id
	where["is_hot"]=is_hot
	where["is_sale"]=is_sale
	list, err := models.GoodsList(offset, 20, where)
	if err != nil{
		util.ResponseWithJson(9002,err,"",c)
		return
	}

	util.ResponseWithJson(200,list,"返回成功",c)
}
