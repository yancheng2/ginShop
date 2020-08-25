package v1

import (
	"fmt"
	"ginShop/models"
	"ginShop/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"reflect"
)

// 分类导航列表接口
func GetCategory(c *gin.Context) {
	category, err := models.GetCategory()
	// 无数据报错
	if gorm.IsRecordNotFoundError(err) {
		util.ResponseWithJson(9006, category, "", c)
		return
	} else if err != nil {
		util.ResponseWithJson(9002, "", "", c)
		return
	}
	fmt.Println(category)
	//分类 数据处理
	var result []models.Category
	tree := CateToTree(category, 0, result)
	util.ResponseWithJson(200, tree, "", c)
}

// 生成分类树
func CateToTree(list []models.Category, parent_id int, result []models.Category) []models.Category {
	var ops []models.Category
	for _, val := range list {
		if val.ParentId == parent_id {
			val.Child = CateToTree(list, val.CatId, ops)
			result = append(result, val)
		}
	}
	return result
}

// 结构体转map
func StructToMapDemo(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
