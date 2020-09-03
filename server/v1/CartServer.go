package v1

import (
	"ginShop/models"
	"ginShop/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)
// 添加商品至购物车
func AddGoodsToCart(c *gin.Context){
	param_goods_id := c.DefaultPostForm("goods_id","")
	param_goods_number := c.DefaultPostForm("goods_number","1")
	// 整理数据
	goods_id,_ := strconv.Atoi(param_goods_id)
	// 数据验证
	vali := validation.Validation{}
	vali.Required(param_goods_id,"goods_id").Message("请选择商品")
	if isOk := checkValidation(&vali, c); isOk == false {
		return
	}
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
	goods_number,_ := strconv.Atoi(param_goods_number)
	goods_price := goodsinfo.GoodsPrice
	goods_sn := goodsinfo.GoodsSn
	saveInfo := map[string]interface{}{"goods_id":goods_id,"goods_number":goods_number,"goods_price":goods_price,"goods_sn":goods_sn}
	// 添加进购物车
	cart, err := models.AddCart("1", 1, saveInfo)
	if err !=nil{
		util.ResponseWithJson(9002,err,"",c)
		return
	}
	util.ResponseWithJson(200,cart,"加购成功",c)
}

// 获取购物车商品信息

func GetCartGoodsList(){

}
