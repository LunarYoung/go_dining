package pkg

//本来用etcd做配置中心，后来用阿里nacos算了

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"go.etcd.io/etcd/client/v3"
//	"time"
//)
//
//type Cig struct {
//	RedisHost       string `json:"redis_host"`
//	RedisDb         int    `json:"redis_db"`
//	BucketName      string `json:"bucketName"`
//	Endpoint        string `json:"endpoint"`
//	AccessKeyId     string `json:"accessKeyId"`
//	AccessKeySecret string `json:"accessKeySecret"`
//	Pre             string `json:"pre"`
//	Es              string `json:"es"`
//	LogPath         string `json:"log_path"`
//	LogName         string `json:"log_name"`
//}
//
//func ETCD() {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"139.159.182.159:2379"}, //如果是集群，就在后面加所有的节点[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
//		DialTimeout: 5 * time.Second,
//	})
//	if err != nil {
//		// handle error!
//		fmt.Printf("connect to etcd failed, err:%v\n", err)
//		return
//	}
//	defer cli.Close()
//
//	//context超时控制
//	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
//	resp, err := cli.Get(ctx, "/config")
//	cancel()
//	if err != nil {
//		fmt.Printf("get from etcd failed,err %v\n", err)
//	}
//
//	//遍历键值对
//
//	//json转为map数据结构
//
//	fmt.Println(len(resp.Kvs))
//	inf := new(Cig)
//	err1 := json.Unmarshal(resp.Kvs[0].Value, inf)
//	if err1 != nil {
//		panic(err)
//	}
//	fmt.Println(inf)
//
//}
//
//func main() {
//
//}
//
////package main
////
////import (
////	"fmt"
////	"github.com/Shopify/sarama"
////	"sync"
////)
////
////消息写入kafka
////func main() {
////	//初始化配置
////	config := sarama.NewConfig()
////	config.Producer.RequiredAcks = sarama.WaitForAll
////	config.Producer.Partitioner = sarama.NewRandomPartitioner
////	config.Producer.Return.Successes = true
////	//生产者
////	client, err := sarama.NewSyncProducer([]string{"139.159.182.159:9092"}, config)
////	if err != nil {
////		fmt.Println("producer close,err:", err)
////		return
////	}
////	defer client.Close()
////
////	for i := 0; i < 5; i++ {
////		//创建消息
////		msg := &sarama.ProducerMessage{}
////		msg.Topic = "cctv1"
////		msg.Value = sarama.StringEncoder("this is a good test,hello kai")
////		//发送消息
////		pid, offset, err := client.SendMessage(msg)
////		if err != nil {
////			fmt.Println("send message failed,", err)
////			return
////		}
////		fmt.Printf("pid:%v offset:%v\n", pid, offset)
////		time.Sleep(time.Second)
////	}
////}
////
////var wg sync.WaitGroup
////
////func main() {
////	consumer, err := sarama.NewConsumer([]string{"139.159.182.159:9092"}, nil)
////	if err != nil {
////		fmt.Println("consumer connect err:", err)
////		return
////	}
////	defer consumer.Close()
////
////	//获取 kafka 主题
////	partitions, err := consumer.Partitions("cctv1")
////	if err != nil {
////		fmt.Println("get partitions failed, err:", err)
////		return
////	}
////
////	for _, p := range partitions {
////		//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
////		partitionConsumer, err := consumer.ConsumePartition("cctv1", p, sarama.OffsetNewest)
////		if err != nil {
////			fmt.Println("partitionConsumer err:", err)
////			continue
////		}
////		wg.Add(1)
////		go func() {
////			for m := range partitionConsumer.Messages() {
////				fmt.Printf("key: %s, text: %s, offset: %d", string(m.Key), string(m.Value), m.Offset)
////			}
////			wg.Done()
////		}()
////	}
////	wg.Wait()
////}
////package main
////
////import (
////"fmt"
////"github.com/Shopify/sarama"
////"time"
////)
////
////func main() {
////	//初始化配置
////	config := sarama.NewConfig()
////	config.Producer.RequiredAcks = sarama.WaitForAll
////	config.Producer.Partitioner = sarama.NewRandomPartitioner
////	config.Producer.Return.Successes = true
////	//生产者
////	client, err := sarama.NewSyncProducer([]string{"139.159.182.159:9092"}, config)
////	if err != nil {
////		fmt.Println("producer close,err:", err)
////		return
////	}
////	defer client.Close()
////
////	for i := 0; i < 5; i++ {
////		//创建消息
////		msg := &sarama.ProducerMessage{}
////		msg.Topic = "cctv1"
////		msg.Value = sarama.StringEncoder("this is a good test,hello kai")
////		//发送消息
////		pid, offset, err := client.SendMessage(msg)
////		if err != nil {
////			fmt.Println("send message failed,", err)
////			return
////		}
////		fmt.Printf("pid:%v offset:%v\n", pid, offset)
////		time.Sleep(time.Second)
////	}
////}
