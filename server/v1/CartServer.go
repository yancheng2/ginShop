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
func AddGoodsToCart(c *gin.Context) {
	param_goods_id := c.DefaultPostForm("goods_id", "")
	param_goods_number := c.DefaultPostForm("goods_number", "1")
	token, exists_t := c.Get("token")
	uid, exists_u := c.Get("ID")
	if exists_t == false || exists_u == false {
		util.ResponseWithJson(9005, "", "", c)
		return
	}
	// 整理数据
	goods_id, _ := strconv.Atoi(param_goods_id)
	// 数据验证
	vali := validation.Validation{}
	vali.Required(param_goods_id, "goods_id").Message("请选择商品")
	if isOk := checkValidation(&vali, c); isOk == false {
		return
	}
	// 查看商品详情  得到价格
	goodsinfo, err := models.GoodsDetails(goods_id)
	if gorm.IsRecordNotFoundError(err) {
		util.ResponseWithJson(1001, "", "", c)
		return
	}
	if err != nil {
		util.ResponseWithJson(9002, "", "", c)
		return
	}
	goods_number, _ := strconv.Atoi(param_goods_number)
	goods_price := goodsinfo.GoodsPrice
	goods_sn := goodsinfo.GoodsSn
	saveInfo := map[string]interface{}{"goods_id": goods_id, "goods_number": goods_number, "goods_price": goods_price, "goods_sn": goods_sn}
	// 添加进购物车
	cart, err := models.AddCart(token.(string), uid.(int), saveInfo)
	if err != nil {
		util.ResponseWithJson(9002, err, "", c)
		return
	}
	util.ResponseWithJson(200, cart, "加购成功", c)
}

// 获取购物车商品信息
func GetCartGoodsList(c *gin.Context) {
	uid, _ := c.Get("ID")

	result, err := models.GetUserCartGoods(uid.(int))
	if models.CheckEmpty(err) {
		util.ResponseWithJson(9002, "", "", c)
		return
	}
	util.ResponseWithJson(200, result, "查询成功", c)
}

// 删除购物车商品
func DelCartGoods(c *gin.Context) {
	// 接值
	id := c.DefaultQuery("cart_id", "")
	cart_id, _ := strconv.Atoi(id)
	uid, _ := c.Get("ID")

	err := models.DelCartGoods(cart_id, uid.(int))
	if err != nil {
		util.ResponseWithJson(9002, err, "", c)
		return
	}
	util.ResponseWithJson(200, "", "删除成功", c)
}
