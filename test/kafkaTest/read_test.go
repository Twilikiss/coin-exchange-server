// Package main
// @Author twilikiss 2024/6/25 10:54:54
package main

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"kafkaTest/database"
	"log"
	"testing"
)

func TestReadOld(t *testing.T) {
	for {
		kafkaConfig := database.KafkaConfig{
			Addr:          "localhost:9092",
			WriteCap:      100,
			ReadCap:       100,
			ConsumerGroup: "test_kafka_consumer_old",
		}
		kafkaCli := database.NewKafkaClient(kafkaConfig)
		kafkaCli.StartReadOld("test_kafka_topic")
		log.Println("正在等待数据传输")
		data := kafkaCli.Read()
		fmt.Println(string(data.Data))
	}
}

// TestReadPlus 测试后发现为了保证数据不会出现积压现象，应该是同一个【Reader】去循环读取数据
func TestReadPlus(t *testing.T) {
	kafkaConfig := database.KafkaConfig{
		Addr:          "localhost:9092",
		WriteCap:      100,
		ReadCap:       100,
		ConsumerGroup: "test_kafka_consumer_plus",
	}
	kafkaCli := database.NewKafkaClient(kafkaConfig)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaCli.C.Addr},
		Topic:    "test_kafka_topic",
		GroupID:  kafkaCli.C.ConsumerGroup,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		//CommitInterval: time.Second, // offset 上报间隔
		//StartOffset: kafka.FirstOffset,
	})
	kafkaCli.StartReadPlus(r)
	for {
		log.Println("正在等待数据传输")
		data := kafkaCli.Read()
		fmt.Println(string(data.Data))
	}
}
