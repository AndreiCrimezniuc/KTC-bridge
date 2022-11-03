package main

import (
	"Portal/Core/internal/initApp"
	"Portal/Core/pkg/ClickhouseServices"
	KafkaHandler "Portal/Core/pkg/KafkaServices"
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/uptrace/go-clickhouse/ch"
	"log"
	"sync"
)

func main() {
	config := initApp.GetConfig()

	chConnection := ClickhouseServices.Connection(&config.CH)

	ClickhouseServices.Ping(&config.CH, chConnection, context.Background())

	wg := sync.WaitGroup{}

	for _, el := range KafkaHandler.GetTopicsFromConfig(&config.Kafka) {
		wg.Add(1)

		config.Kafka.Consumer = el
		consumer := KafkaHandler.Consumer(el.Topic+"-"+el.Type, config.Kafka.GroupId)

		go handleTopic(consumer, chConnection, config)
	}

	wg.Wait()

	defer func() {
		er := chConnection.Close()
		if er != nil {
			log.Fatalln("Can't close Clickhouse connection")
		}
	}()
}

func handleTopic(consumer *kafka.Reader, chConnection *ch.DB, config *initApp.Config) {
	defer KafkaHandler.CloseConnection(consumer)
	for {
		chunk, messagesChunk := KafkaHandler.FetchChunk(consumer, config.Kafka)

		err := ClickhouseServices.Write(chunk, context.Background(), chConnection, config.CH)

		if err == nil {
			KafkaHandler.CommitChunk(consumer, messagesChunk, config.Kafka)
		}
	}
}
