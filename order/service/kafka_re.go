package service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func KafkaRe() {
	fmt.Println("kafka start")
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{"106.12.108.5:9092"}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	partitionList, err := consumer.Partitions("task-status-data") // 通过topic获取到所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition: ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList { // 遍历所有的分区
		pc, err := consumer.ConsumePartition("task", int32(partition), sarama.OffsetNewest) // 针对每个分区创建一个分区消费者
		if err != nil {
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { // 为每个分区开一个go协程取值
			for msg := range pc.Messages() { // 阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
