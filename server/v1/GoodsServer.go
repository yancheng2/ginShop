package v1

import (
	"fmt"
	"ginShop/models"
	"ginShop/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
	"sync"
)

// 商品列表+搜索
func GetGoodsList(c *gin.Context) {
	page_s := c.Query("page")
	cat_id := c.Query("cat_id")
	keyword := c.Query("keyword")
	is_hot := c.Query("is_hot")
	is_sale := c.Query("is_sale")

	if page_s == "" {
		page_s = "1"
	}
	page, _ := strconv.Atoi(page_s)
	limit := 20
	offset := (page - 1) * 20

	where := make(map[string]interface{}) //查询条件
	cat := make(map[string]interface{})   //filter cat_id
	hot := make(map[string]interface{})   //filter is_hot
	sale := make(map[string]interface{})  //filter is_hot
	cat["field"] = "cat_id"
	cat["val"] = cat_id
	hot["field"] = "is_hot"
	hot["val"] = is_hot
	sale["field"] = "is_sale"
	sale["val"] = is_sale
	where["limit"] = limit
	where["offset"] = offset
	where["query"] = keyword
	where["filter"] = map[int]interface{}{0: cat, 1: hot, 2: sale}

	result, err := util.Search("go_goods", where)

	if err != nil {
		util.ResponseWithJson(9002, err, "", c)
		return
	}
	// 对es搜索后的数据进行处理整合
	total, list := util.SearchResultHandle(result)
	ret := map[string]interface{}{"total": total, "list": list}

	util.ResponseWithJson(200, ret, "返回成功", c)
}

// 商品详情
func GoodsDetails(c *gin.Context) {
	goods_id := c.DefaultQuery("goods_id", "")
	atoi, _ := strconv.Atoi(goods_id)
	details, err := models.GoodsDetails(atoi)
	if gorm.IsRecordNotFoundError(err) { //如果没有错误只是无数据则注册
		util.ResponseWithJson(9006, err, "", c)
		return
	}
	if err != nil {
		util.ResponseWithJson(9002, err, "", c)
		return
	}
	util.ResponseWithJson(200, details, "查询成功", c)
}

// 暂时-创建索引
func CreateIndex(c *gin.Context) {
	body := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"goods_list":{
				"properties":{
					"goods_id":{"type":"integer"},
					"cat_id":{"type":"integer"},
					"goods_sn":{"type":"keyword"},
					"goods_name":{"type":"text"},
					"goods_img":{"type":"text"},
					"goods_number":{"type":"integer"},
					"sort_order":{"type":"integer"},
					"goods_price":{"type":"float"},
					"is_hot":{"type":"boolean"},
					"is_sale":{"type":"boolean"}
					"is_new":{"type":"boolean"}
					"create_time":{"type":"text"}
				}
			}
		}
	}`
	util.CreateIndex("go_goods", body)
}

// 暂时-同步goods表到es
func SyncGoodsToEs(c *gin.Context) {
	count, err := models.GoodsCount()
	if err != nil {
		util.ResponseWithJson(9002, err, "", c)
		return
	}
	limit := 100
	pageCount := math.Ceil(float64(count) / float64(limit))
	wg := sync.WaitGroup{}
	for i := 1; i <= int(pageCount); i++ {
		wg.Add(1)
		go syncEs(i, limit, &wg)
	}
	wg.Wait()
}

func syncEs(page, limit int, wg *sync.WaitGroup) {
	offset := (page - 1) * limit
	list, err := models.GoodsList(offset, limit, nil)
	if err != nil {
		//return err
	}

	util.BatchAddDocument("go_goods", "goods_list", list)
	fmt.Println("同步成功了-i")
	wg.Done()
}
