package models

import (
	"fmt"
	"time"
)

type Cart struct {
	CartId     int     `gorm:"type:int(11);not null;primaryKey;autoIncrement;"`
	UserId     int     `gorm:"type:int(11);"`
	GoodsId    int     `gorm:"type:int(11);not null;"`
	GoodsSn    string  `gorm:"type:varchar(30);not null;"`
	GoodsPrice float64 `gorm:"type:float(10);not null;DEFAULT:0.00;"`
	CartNumber int     `gorm:"type:int(11);not null;default:0;"`
	Token      string  `gorm:"type:varchar(200);"`
	Status     int     `gorm:"type:tinyint(1);not null;default:0;"`
	CreateTime string  `gorm:"type:datetime;not null;"`
}

type CartList struct {
	Cart
	Goods
}

func AddCart(Token string, user_id int, goodsData map[string]interface{}) (Cart, error) {
	if goodsData == nil {

	}
	var info Cart
	info.GoodsId = goodsData["goods_id"].(int)
	info.CartNumber = goodsData["goods_number"].(int)
	info.GoodsPrice = goodsData["goods_price"].(float64)
	info.GoodsSn = goodsData["goods_sn"].(string)
	info.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	info.Token = Token
	info.UserId = user_id
	err := db.Omit("CatId", "Status").Create(&info).Error

	return info, err
}

// 获取用户的购物车列表
func GetUserCartGoods(uid int) ([]CartList, error) {
	var list []CartList
	err := db.Table("sh_cart as c").Select("c.cart_id,c.cart_number,g.*").Joins("left join sh_goods as g on g.goods_id = c.goods_id").Where("user_id=?", uid).Where("c.status=?", 0).Find(&list).Error
	return list, err
}

func DelCartGoods(cart_id, uid int) error {
	var info Cart
	info.Status = 1
	err := db.Table("sh_cart").Where("cart_id=?", cart_id).Where("user_id=?", uid).Update(info).Error
	return err
}

// 批量新增-暂时未成功
func BatchAddCart(Token string, user_id int, goodsData map[int]map[string]interface{}) ([]Cart, error) {
	if goodsData == nil {

	}
	var data []Cart
	var info Cart
	for _, val := range goodsData {
		info.GoodsId = val["goods_id"].(int)
		info.CartNumber = val["goods_number"].(int)
		info.GoodsPrice = val["goods_price"].(float64)
		info.GoodsSn = val["goods_sn"].(string)
		info.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		info.Token = Token
		info.UserId = user_id
		data = append(data, info)
	}
	fmt.Println("info:", info)
	fmt.Println("data:", data)
	err := db.Omit("CatId", "Status").Create(&info).Error

	if err != nil {
		return nil, err
	}
	return data, err
}
