package KafkaHandler

import (
	"Portal/Core/pkg/ClickhouseServices"
	"context"
	"github.com/segmentio/kafka-go"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Probes                  map[string]string   `yaml:"probes,omitempty"`
	Files                   map[string][]string `yaml:"files"`
	GroupId                 string              `yaml:"groupId"`
	AutoOffset              int64               `yaml:"autoOffset"`
	ChunkSize               int                 `yaml:"chunkSize"`
	TimeToFetch             time.Duration       `yaml:"timeToFetch"`
	RetryBeforeWriteTheTail uint                `yaml:"retryBeforeWriteTheTail"`
	Consumer                ConsumerData
}

func Consumer(topic, groupId string) *kafka.Reader {
	c := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KF_SERVER_ADDRESS")},
		Topic:   topic,
		GroupID: groupId,
	})
	return c
}

func FetchChunk(consumer *kafka.Reader, config Config) ([]interface{}, []kafka.Message) {
	chunk := make([]interface{}, 0, config.ChunkSize)
	messageChunk := make([]kafka.Message, 0, config.ChunkSize)

	var (
		message kafka.Message
		err     error
		tries   uint
	)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), config.TimeToFetch)

		message, err = consumer.FetchMessage(ctx)
		defer cancel()

		if err == nil {
			el, er := ClickhouseServices.ConvertDataToCHModel(message.Topic, string(message.Value), string(message.Key), config.Probes, config.Consumer.Type)

			if er != nil {
				log.Println(er)
				continue
			}

			chunk = append(chunk, el)
			messageChunk = append(messageChunk, message)
			tries = 0
		} else {
			log.Printf("Error in reading from kafka topic - %s. %s. Trying again...", consumer.Config().Topic, err)

			if len(chunk) != 0 {
				tries++
			}
		}

		if len(chunk) == config.ChunkSize || (tries >= config.RetryBeforeWriteTheTail && len(chunk) > 0) {
			return chunk, messageChunk
		}
	}
}

func CommitChunk(consumer *kafka.Reader, messages []kafka.Message, config Config) {
	ctx, cancel := context.WithTimeout(context.Background(), config.TimeToFetch)

	er := consumer.CommitMessages(ctx, messages...)
	defer cancel()

	if er != nil {
		log.Fatalf("Error in commiting in kafka topic - %s. %s. Trying again...\n", consumer.Config().Topic, er)
	}
}

func GetTopicsFromConfig(c *Config) (grpTopics []ConsumerData) {
	for in, el := range c.Files {
		for _, typeConn := range el {
			grpTopics = append(grpTopics, ConsumerData{Topic: in, Type: typeConn})
		}
	}
	return
}

func CloseConnection(consumer *kafka.Reader) {
	err := consumer.Close()
	if err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func GetConfig() *Config {
	fh, err := os.Open("././configs/kafka_config.yaml")

	if err != nil {
		log.Fatalln("can't open kafka config in 'configs/kafka_config.yaml'")
	}

	kafkaConfig := Config{}

	if err = yaml.NewDecoder(fh).Decode(&kafkaConfig); err != nil {
		log.Fatalln(err)
	}

	return &kafkaConfig
}

type ConsumerData struct {
	Topic string
	Type  string
}
