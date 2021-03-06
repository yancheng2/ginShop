package models

type Goods struct {
	GoodsId     int     `gorm:"type:int(11);not null;PRIMARY_KEY:true;AUTO_INCREMENT:true;"`
	GoodsName   string  `gorm:"type:varchar(100);not null;"`
	GoodsSn   	string  `gorm:"type:varchar(60);not null;"`
	GoodsPrice  float64 `gorm:"type:float(10);not null;DEFAULT:0.00;"`
	GoodsNumber int     `gorm:"type:int(11);not null;default:0;"`
	CatId       int     `gorm:"type:int(11);not null;"`
	GoodsImg    string  `gorm:"type:varchar(100);not null;"`
	IsHot       int     `gorm:"type:tinyint(1);not null;default:0;"`
	IsSale      int     `gorm:"type:tinyint(1);not null;default:0;"`
	IsNew       int     `gorm:"type:tinyint(1);not null;default:0;"`
	SortOrder   int     `gorm:"type:smallint(4);not null;default:100;"`
	CreateTime  string  `gorm:"type:datetime;not null;"`
	CatName     string
}

// 商品列表
func GoodsList(offset, limit int, where map[string]interface{}) ([]Goods, error) {
	var list []Goods
	err := db.Offset(offset).Limit(limit).Find(&list).Error

	return list, err
}

// 商品总数
func GoodsCount() (int, error) {
	var count int
	err := db.Model(&Goods{}).Count(&count).Error
	return count, err
}

// 商品详情
func GoodsDetails(goods_id int) (Goods, error) {
	var info Goods
	err := db.Select("sh_goods.*,c.cat_name").Joins("left join sh_category as c on c.cat_id = sh_goods.cat_id").Where("goods_id=?", goods_id).First(&info).Error
	return info, err
}
