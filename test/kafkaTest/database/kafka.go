package database

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
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
	C         KafkaConfig
	closed    bool
	mutex     sync.Mutex
}

func NewKafkaClient(c KafkaConfig) *KafkaClient {
	return &KafkaClient{
		C: c,
	}
}

func (k *KafkaClient) StartWrite() {
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.C.Addr),
		Balancer: &kafka.LeastBytes{},
	}
	k.w = w
	k.writeChan = make(chan KafkaData, k.C.WriteCap)
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

// SendSync 同步发送到Kafka
func (k *KafkaClient) SendSync(data KafkaData) error {
	w := &kafka.Writer{
		Addr:     kafka.TCP(k.C.Addr),
		Balancer: &kafka.LeastBytes{},
	}
	messages := []kafka.Message{
		{
			Topic: data.Topic,
			Key:   data.Key,
			Value: data.Data,
		},
	}
	w.AllowAutoTopicCreation = true
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = w.WriteMessages(ctx, messages...)
	return err
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
					time.Sleep(time.Millisecond * 250)
					success = false
					continue
				}
				if err != nil {
					success = false
					log.Printf("kafka send writemessage err %s \n", err.Error())
				}
			}
			if !success {
				//重新放进去等待消费
				k.Send(data)
			}
		}
	}

}

func (k *KafkaClient) StartReadPlus(r *kafka.Reader) {
	k.r = r
	k.readChan = make(chan KafkaData, k.C.ReadCap)
	go k.readMsg()
}

func (k *KafkaClient) StartReadOld(topic string) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.C.Addr},
		Topic:    topic,
		GroupID:  k.C.ConsumerGroup,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	k.r = r
	k.readChan = make(chan KafkaData, k.C.ReadCap)
	go k.readMsg()
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

func (k *KafkaClient) RPut(data KafkaData) {
	k.readChan <- data
}
