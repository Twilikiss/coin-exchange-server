// Package database
// @Author twilikiss 2024/5/8 16:51:51
package database

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type KafkaConfig struct {
	Addr          string `json:"addr,optional"`
	WriteCap      int    `json:"writeCap,optional"`
	ReadCap       int    `json:"readCap,optional"`
	ConsumerGroup string `json:"ConsumerGroup,optional"`
}
type KafkaData struct {
	Topic string
	Key   []byte
	Data  []byte
}
type KafkaClient struct {
	w         *kafka.Writer
	r         *kafka.Reader
	readChan  chan KafkaData
	writeChan chan KafkaData
	c         KafkaConfig
	closed    bool
	mutex     sync.Mutex
}

func NewKafkaClient(c KafkaConfig) *KafkaClient {
	return &KafkaClient{
		c: c,
	}
}

func (k *KafkaClient) StartWrite() {
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.c.Addr),
		Balancer: &kafka.LeastBytes{},
	}
	k.w = w
	k.writeChan = make(chan KafkaData, k.c.WriteCap)
	go k.sendKafka()
}

func (k *KafkaClient) Send(data KafkaData) {
	defer func() {
		if err := recover(); err != nil {
			k.closed = true
		}
	}()
	k.writeChan <- data
	k.closed = false
}

func (k *KafkaClient) Close() {
	if k.w != nil {
		k.w.Close()
		k.mutex.Lock()
		defer k.mutex.Unlock()
		if !k.closed {
			close(k.writeChan)
			k.closed = true
		}
	}
	if k.r != nil {
		k.r.Close()
	}
}

// 等待writeChan的数据到来
func (k *KafkaClient) sendKafka() {
	for {
		select {
		case data := <-k.writeChan:
			messages := []kafka.Message{
				{
					Topic: data.Topic,
					Key:   data.Key,
					Value: data.Data,
				},
			}
			var err error
			const retries = 3
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			success := false
			for i := 0; i < retries; i++ {
				// attempt to create topic prior to publishing the message
				err = k.w.WriteMessages(ctx, messages...)
				if err == nil {
					success = true
					break
				}
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					// 触发重试机制，需要等待250ms后再次重试
					logx.Infof("触发超时重试，原因为error=%s", err.Error())
					time.Sleep(time.Millisecond * 250)
					success = false
					continue
				}
				if err != nil {
					success = false
					logx.Errorf("kafka send writemessage err %s \n", err.Error())
				}
			}
			if !success {
				//重新放进去等待消费
				k.Send(data)
			}
		}
	}

}

func (k *KafkaClient) StartRead(topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.c.Addr},
		Topic:    topic,
		GroupID:  k.c.ConsumerGroup,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	k.r = r
	k.readChan = make(chan KafkaData, k.c.ReadCap)
	go k.readMsg()
}

// StartReadNew 开启kafka的读取数据的功能
// 这里的代码和其他模块下的同名方法不一样，原因是exchange模块里面有超过两个不同的topic，他们都混用者同一个client，会出现读取的数据混乱
func (k *KafkaClient) StartReadNew(topic string) *KafkaClient {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.c.Addr},
		Topic:    topic,
		GroupID:  k.c.ConsumerGroup,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	client := NewKafkaClient(k.c)
	client.r = r
	client.readChan = make(chan KafkaData, k.c.ReadCap)
	go client.readMsg()
	return client
}

func (k *KafkaClient) readMsg() {
	for {
		m, err := k.r.ReadMessage(context.Background())
		if err != nil {
			logx.Error(err)
			continue
		}
		data := KafkaData{
			Key:  m.Key,
			Data: m.Value,
		}
		k.readChan <- data
	}
}

func (k *KafkaClient) Read() KafkaData {
	msg := <-k.readChan
	return msg
}
