package consumer

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "mygroup",
		"auto.offset.reset": "earliest",
	}
	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)", err, msg)
		}
	}

	c.Close()
}
