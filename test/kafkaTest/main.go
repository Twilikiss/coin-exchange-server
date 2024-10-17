// Package kafkaTest
// @Author twilikiss 2024/6/25 9:59:59
package main

import (
	"fmt"
	"kafkaTest/database"
	"log"
	"time"
)

func main() {
	kafkaConfig := database.KafkaConfig{
		Addr:          "localhost:9092",
		WriteCap:      100,
		ReadCap:       100,
		ConsumerGroup: "test_kafka_consumer",
	}
	kafkaCli := database.NewKafkaClient(kafkaConfig)
	index := 0
	for {
		kafkaData := database.KafkaData{
			Topic: "test_kafka_topic",
			Key:   []byte("Elysia"),
			Data:  []byte(fmt.Sprintf("index = %d ,Elysia is best!", index)),
		}
		err := kafkaCli.SendSync(kafkaData)
		if err != nil {
			fmt.Println(err)
		}
		log.Println("消息已发送,index=", index)
		index++
		time.Sleep(1 * time.Second)
	}

}
