package main

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost",
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
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

	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kakfa", "Golang", "client"} {
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}
		p.Produce(message, nil)
	}

	p.Flush(15 * 1000)

}
