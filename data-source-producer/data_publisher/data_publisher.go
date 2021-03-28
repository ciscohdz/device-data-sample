package data_publisher

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type DeviceData struct {
	DeviceId string
	Value    float32
	Ts       time.Time
}

func Publish(dataChan <-chan DeviceData) {

	cfg := kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"compression.type":  "gzip",
	}

	p, err := kafka.NewProducer(&cfg)
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// goroutine which handles the events on the producer
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// not publish the records
	topic := "raw-device-signals"
	topicPartition := kafka.TopicPartition{Topic: &topic,
		Partition: kafka.PartitionAny}

	for result := range dataChan {
		payload, err := toBytes(&result)
		if err == nil {
			msg := kafka.Message{
				TopicPartition: topicPartition,
				Value:          payload,
			}
			p.Produce(&msg, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

}

func toBytes(dd *DeviceData) ([]byte, error) {
	bits, err := json.Marshal(dd)
	if err != nil {
		return nil, err
	}
	return bits, nil
}
