package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

/*
	Initialize NewConfig configuration sarama.NewConfig
	Create producer sarama.NewSyncProducer
	Create message sarama.ProducerMessage
	Send message client.SendMessage
*/
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "TestTopic"
	msg.Value = sarama.StringEncoder("this is a test")

	client, err := sarama.NewSyncProducer([]string{"172.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed,", err)
		return
	}

	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
