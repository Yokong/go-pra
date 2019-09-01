package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

type MsgBody struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx.log"

	mb := MsgBody{
		Type:    "test",
		Name:    "Yoko",
		Content: "this is a test message with message body",
	}
	mbBytes, err := json.Marshal(mb)
	if err != nil {
		fmt.Println("Marshal MsgBody err:", err)
		return
	}
	msg.Value = sarama.ByteEncoder(mbBytes)
	client, err := sarama.NewSyncProducer([]string{"192.168.56.20:9092"}, config)
	if err != nil {
		fmt.Println("producer close err: ", err)
		return
	}
	defer client.Close()

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed: ", err)
		return
	}
	fmt.Println("pid: ", pid, " offset: ", offset)
}
