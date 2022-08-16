package pkg

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type ContractChatMsgInfo struct {
	OrgId      string `json:"org_id"`
	ConType    int    `json:"con_type"`
	SendFrom   string `json:"send_from" `   //
	SendTo     string `json:"send_to"`      //
	MsgContent string `json:"msg_content" ` // 内容
	MsgDate    string `json:"msg_date"`     // 创建时间
}

var db *mongo.Database           // database 话柄
var collection *mongo.Collection // collection 话柄

// ConnectToDB 1、pool连接池模式
func ConnectToDB() {
	// 设置连接超时时间
	uri := RemoteViper.GetString("mongo")

	name := "msg"               //数据库名
	timeout := time.Duration(2) // 链接超时时间

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)

	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)

	}
	// 返回 client
	db = client.Database(name)

}

// AddOne 2、新增一条数据
func AddOne(t interface{}, table string) {

	_, err := db.Collection(table).InsertOne(context.TODO(), &t)
	if err != nil {
		log.Println(err)
		return
	}
}

// GetList 获取多条数据

func GetList(m bson.M, size int64) (rep []ContractChatMsgInfo) {
	collection = db.Collection("msg")
	findOptions := options.Find()
	findOptions.SetLimit(size)
	cur, err := collection.Find(context.Background(), m, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	err = cur.All(context.Background(), &rep)
	if err != nil {
		log.Fatal(err)
	}
	_ = cur.Close(context.Background())
	return rep
}

//、统计collection的数据总数
//func Count() {
//	count, err := collection.CountDocuments(context.Background(), bson.D{})
//	if err != nil {
//		log.Fatal(count)
//	}
//	log.Println("collection.CountDocuments:", count)
//}
