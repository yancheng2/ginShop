package util

import (
	"fmt"
	"ginShop/models"
	"ginShop/pkg/setting"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"time"
)

var es *elastic.Client

//连接初始化
func inits() error {
	var err error
	es, err = elastic.NewClient(
		elastic.SetURL(setting.ElasticsearchSetting.Host),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(setting.ElasticsearchSetting.SetHealthcheckInterval*time.Second),
		elastic.SetGzip(setting.ElasticsearchSetting.SetGzip),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	return err
}

// 创建索引
func CreateIndex(indexName, body string) (*elastic.IndicesCreateResult, error) {
	err := inits()
	if err != nil {
		return nil, err
	}
	var ctx = context.Background()
	result, err := es.CreateIndex(indexName).BodyString(body).Do(ctx)
	return result, err
}

// 批量新增文档
//func BatchAddDocument(index_name, type_name string, data interface{}) (*elastic.BulkResponse, error) {
func BatchAddDocument(index_name, type_name string, data []models.Goods) (*elastic.BulkResponse, error) {
	err := inits()
	if err != nil {
		return nil, err
	}
	bulk := es.Bulk()
	//switch items := data.(type){
	//	case []models.Goods:
	//	default:return nil,err
	//}
	for _, val := range data {
		request := elastic.NewBulkIndexRequest().Index(index_name).Type(type_name).Doc(val)
		bulk.Add(request)
	}
	result, err := bulk.Do(context.Background())
	return result, err
}

// 查询文档
func GetDocument(indexName, typeName, id string) (*elastic.IndexResponse, error) {
	err := inits()
	if err != nil {
		return nil, err
	}
	res, err := es.Index().Index(indexName).Type(typeName).Id(id).Do(context.Background())
	return res, err
}

//搜索
func Search(indexName string, where map[string]interface{}) (*elastic.SearchResult, error) {
	err := inits()
	if err != nil {
		return nil, err
	}

	search := es.Search(indexName).Index(indexName)

	// 过滤
	boolQuery := elastic.NewBoolQuery()
	if _, ok := where["filter"]; ok {
		filterW := where["filter"].(map[int]interface{})
		var querys []elastic.Query
		fmt.Println("filter:", filterW)
		for _, val := range filterW {
			if val != nil {
				v := val.(map[string]interface{})
				fmt.Println("v:", v)
				if v["val"] != "" {
					field := v["field"].(string)
					val := v["val"]
					fmt.Println("field:", field, "val:", val)
					filter := elastic.NewTermQuery(field, val)
					if filter == nil {
						continue
					}
					fmt.Println("NewTermQuery后：", filter)
					querys = append(querys, filter)
				}
			}
		}
		if querys != nil {
			fmt.Println("filter querys:", querys)
			boolQuery.Filter(querys...)
		}
	}
	// 查询 模糊
	if _, ok := where["query"]; ok {
		queryW := where["query"].(string)
		if queryW != "" {
			fmt.Println(queryW)
			query := elastic.NewMultiMatchQuery(queryW)
			boolQuery.Must(query)
		}
	}
	if boolQuery != nil {
		search.Query(boolQuery)
	}

	// 排序
	if _, ok := where["sort"]; ok {
		field := where["sort"].(map[string]interface{})["field"].(string)
		val := where["sort"].(map[string]interface{})["val"].(bool)
		search.Sort(field, val)
	}

	if _, ok := where["offset"]; ok {
		offset := where["offset"].(int)
		search.From(offset)
	}

	if _, ok := where["limit"]; ok {
		limit := where["limit"].(int)
		search.Size(limit)
	}

	if _, ok := where["pretty"]; ok {
		pretty := where["pretty"].(bool)
		search.Pretty(pretty)
	}

	result, err := search.Do(context.Background())
	return result, err

}

// 对搜索结果进行整合处理
func SearchResultHandle(result *elastic.SearchResult) (int64, []interface{}) {
	var list []interface{}
	total := result.Hits.TotalHits // 数据总数量
	hits := result.Hits.Hits       //结果
	for _, val := range hits {
		list = append(list, val.Source)
	}
	return total, list
}
