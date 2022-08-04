package pkg

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"order/model"
	"reflect"
)

var client *elastic.Client
var host = "http://106.12.108.5:19200"

func init() {

	options := []elastic.ClientOptionFunc{
		elastic.SetURL("http://106.12.108.5:19200"), elastic.SetSniff(false),
	}
	var err error
	client, err = elastic.NewClient(options...)
	if err != nil {
		panic(err)
	}

	//client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL("http://106.12.108.5:19200"))
	//if err != nil {
	//	fmt.Printf(err.Error())
	//	panic(err)
	//
	//}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

/*下面是简单的CURD*/

//创建
func Create() {

	//使用结构体
	e1 := model.Food{Name: "Jane", Price: "23.4", Total: 343, Pic: []string{"www.", "ssss"}}
	put1, err := client.Index().
		Index("food").
		Id("3").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

}

//删除
func Delete() {

	res, err := client.Delete().Index("megacorp").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func Update() {
	res, err := client.Update().
		Index("food").
		Id("3").
		Doc(map[string]interface{}{"Name": "曾秦刚"}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

////查找
//func Gets() {
//	//通过id查找
//	get1, err := client.Get().Index("megacorp").Id("2").Do(context.Background())
//	if err != nil {
//		panic(err)
//	}
//	if get1.Found {
//		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
//	}
//}

//搜索
func Query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	//res, err = client.Search("food").Do(context.Background())
	//PrintEmployee(res, err)
	//字段相等
	//q := elastic.NewQueryStringQuery("Name:Jane")
	//res, err = client.Search("food").Query(q).Do(context.Background())
	//if err != nil {
	//	println(err.Error())
	//}
	//PrintEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	//boolQ := elastic.NewBoolQuery()
	//boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	//boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	//res, err = client.Search("megacorp").Query(q).Do(context.Background())
	//PrintEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("Name", "秦")
	res, err = client.Search("food").Query(matchPhraseQuery).Do(context.Background())
	PrintEmployee(res, err)

	////分析 interests
	//aggs := elastic.NewTermsAggregation().Field("interests")
	//res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	//PrintEmployee(res, err)

}

//简单分页
func List(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("food").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	PrintEmployee(res, err)

}

//打印查询到的Employee
func PrintEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ model.Food
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(model.Food)
		fmt.Printf("%#v\n", t)
	}
}
