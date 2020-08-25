package models

type Category struct {
	DB
	Child []Category
}

// 数据表结构
type DB struct {
	CatId    int    `gorm:"type:int(11);not null;PRIMARY_KEY:true;AUTO_INCREMENT:true;"`
	CatName  string `gorm:"type:varchar(100);not null;"`
	ParentId int    `gorm:"type:int(11);default:0;"`
	Path     string `gorm:"type:varchar(20);"`
}

// 获取分类列表
func GetCategory() ([]Category, error) {
	var list []Category
	err := db.Find(&list).Error
	return list, err
}
