package v1

import (
	"ginShop/models"
	"ginShop/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

func AddGoodsToCart(c *gin.Context){
	param_goods_id := c.DefaultPostForm("goods_id","")
	param_goods_number := c.DefaultPostForm("goods_number","1")

	vali := validation.Validation{}
	vali.Required(param_goods_id,"goods_id").Message("请选择商品")
	if isOk := checkValidation(&vali, c); isOk == false {
		return
	}
	goods_id,_ := strconv.Atoi(param_goods_id)
	goods_number,_ := strconv.Atoi(param_goods_number)

	// 查看商品详情  得到价格
	goodsinfo,err := models.GoodsDetails(goods_id)
	if gorm.IsRecordNotFoundError(err){
		util.ResponseWithJson(1001,"","",c)
		return
	}
	if err!= nil{
		util.ResponseWithJson(9002,"","",c)
		return
	}
	goods_price := goodsinfo.GoodsPrice
	goods_sn := goodsinfo.GoodsSn
	saveInfo := map[string]interface{}{"goods_id":goods_id,"goods_number":goods_number,"goods_price":goods_price,"goods_sn":goods_sn}

	saveData := map[int]map[string]interface{}{0:saveInfo}

	cart, err := models.AddCart("1", 1, saveData)
	if err !=nil{
		util.ResponseWithJson(9002,err,"",c)
		return
	}
	util.ResponseWithJson(200,cart,"加购成功",c)
}
