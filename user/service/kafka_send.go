package service

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func KafkaSend() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follower都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，我们默认设置32个分区
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "task"
	msg.Value = sarama.StringEncoder("producer kafka messages...")

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"106.12.108.5:9092"}, config)
	if err != nil {
		fmt.Println("Producer closed, err:", err)
		return
	}
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
