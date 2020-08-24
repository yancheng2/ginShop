package models

type Goods struct {
	GoodsId 		int			`gorm:"type:int(11);not null;PRIMARY_KEY:true;AUTO_INCREMENT:true;"`
	GoodsName		string		`gorm:"type:varchar(100);not null;"`
	GoodsPrice		float64		`gorm:"type:float(10);not null;DEFAULT:0.00;"`
	GoodsNumber		int			`gorm:"type:int(11);not null;default:0;"`
	CatId			int			`gorm:"type:int(11);not null;"`
	GoodsImg		string		`gorm:"type:varchar(100);not null;"`
	IsHot			int			`gorm:"type:tinyint(1);not null;default:0;"`
	IsSale			int			`gorm:"type:tinyint(1);not null;default:0;"`
	CreateTime		string		`gorm:"type:datetime;not null;"`
}

// 商品列表
func GoodsList(offset,limit int,where map[string]interface{})(Goods,error){
	var list Goods
	err := db.Offset(offset).Limit(limit).Find(&list).Error

	return list,err
}