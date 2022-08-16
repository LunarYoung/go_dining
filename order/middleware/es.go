package middleware

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"order/model"
	"order/model/req"
	"reflect"
	"time"
)

var client *elastic.Client

func InitEs() {

	options := []elastic.ClientOptionFunc{
		elastic.SetURL(RemoteViper.GetString("es")), elastic.SetSniff(false),
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
	info, code, err := client.Ping(RemoteViper.GetString("es")).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(RemoteViper.GetString("es"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

//创建
func Create(m interface{}, i string, id string) {
	put1, err := client.Index().
		Index(i).
		Id(id).
		BodyJson(m).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

}

// Update
// @description: 修改订单状态
// @param r
// @param i
// @2022-08-06 16:35:30
func Update(r req.OrderSearchReq, i string) {
	res, err := client.Update().
		Index(i).
		Id(r.OrderId).
		Doc(map[string]interface{}{"status": r.Status}).
		Do(context.Background())
	if err != nil {
		//es有乐观锁，休眠一秒再插入，重试
		time.Sleep(1 * time.Second)
		_, err1 := client.Update().
			Index(i).
			Id(r.OrderId).
			Doc(map[string]interface{}{"status": r.Status}).
			Do(context.Background())
		if err1 != nil {
			//logrus.Info("第二次插入失败")
			println(err.Error())
		}

	}
	fmt.Printf("update age %s\n", res.Result)
}

func UpdateMenu(r req.MenuChangeReq, i string) {

	var m = make(map[string]interface{})
	m["uuid"] = r.Uuid

	if r.Name != "" {
		m["name"] = r.Name
	}
	if r.Status != 0 {
		m["status"] = r.Status
	}
	if r.Price != 0 {
		m["price"] = r.Price
	}
	if r.MaxSell != 0 {
		m["max_sell"] = r.MaxSell
	}
	_, err := client.Update().
		Index(i).
		Id(r.Uuid).
		Doc(m).
		Do(context.Background())
	if err != nil {

		println(err.Error())
	}

}

// Query
// @description: 查询订单
// @param m
// @param i
// @return re
// @return count
// @2022-08-06 17:30:18
func Query(m req.OrderSearchReq, i string) (re []model.Order, count int64) {

	boolQuery := elastic.NewBoolQuery().Must()
	termsQuery1 := elastic.NewMatchPhraseQuery("org_id", m.OrgId)
	boolQuery.Must(termsQuery1)
	if m.Name != "" {
		termsQuery2 := elastic.NewMatchPhraseQuery("name", m.Name)
		boolQuery.Must(termsQuery2)
	}

	if m.Time != "" {
		termsQuery3 := elastic.NewMatchPhraseQuery("time", m.Time)
		boolQuery.Must(termsQuery3)
	}
	if m.Status != 0 {
		termsQuery4 := elastic.NewMatchPhraseQuery("status", m.Status)
		boolQuery.Must(termsQuery4)
	}

	if m.OrderId != "" {
		termsQuery5 := elastic.NewMatchPhraseQuery("order_id", m.OrderId)
		boolQuery.Must(termsQuery5)
	}

	res, err := client.Search(i).Query(boolQuery).
		Size(int(m.Page.PageSize)).
		From(int(m.Page.PageIndex*m.Page.PageSize - m.Page.PageSize)).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	count, er := client.Count(i).Query(boolQuery).Do(context.Background())
	if er != nil {
		println(err.Error())
	}

	return ToEntity(res, err), count
}

func QueryMenu(m req.MenuSearchReq, i string) (re []model.Menu, count int64) {

	boolQuery := elastic.NewBoolQuery().Must()
	termsQuery1 := elastic.NewMatchPhraseQuery("org_id", m.OrgId)
	boolQuery.Must(termsQuery1)

	if m.Score != 0 {
		termsQuery2 := elastic.NewMatchPhraseQuery("score", m.Score)
		boolQuery.Must(termsQuery2)
	}

	if m.Name != "" {
		termsQuery3 := elastic.NewMatchPhraseQuery("name", m.Name)
		boolQuery.Must(termsQuery3)
	}
	if m.Price != 0 {
		termsQuery4 := elastic.NewMatchPhraseQuery("price", m.Price)
		boolQuery.Must(termsQuery4)
	}

	if m.SalesVolumeMonth != 0 {
		termsQuery5 := elastic.NewMatchPhraseQuery("sales_volume_month", m.SalesVolumeMonth)
		boolQuery.Must(termsQuery5)
	}

	if m.Classification != "" {
		termsQuery6 := elastic.NewMatchPhraseQuery("classification", m.Classification)
		boolQuery.Must(termsQuery6)
	}

	res, err := client.Search(i).Query(boolQuery).
		Size(int(m.Page.PageSize)).
		From(int(m.Page.PageIndex*m.Page.PageSize - m.Page.PageSize)).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	count, er := client.Count(i).Query(boolQuery).Do(context.Background())
	if er != nil {
		println(err.Error())
	}

	return ToEntityMenu(res, err), count
}

func ToEntityMenu(res *elastic.SearchResult, err error) (re []model.Menu) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ model.Menu
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(model.Menu)
		re = append(re, t)
	}
	return re
}

func ToEntity(res *elastic.SearchResult, err error) (re []model.Order) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ model.Order
	for _, item := range res.Each(reflect.TypeOf(typ)) {
		t := item.(model.Order)
		re = append(re, t)
	}
	return re
}
