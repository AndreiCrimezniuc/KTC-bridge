package initApp

import (
	"Portal/Core/pkg/ClickhouseServices"
	KafkaHandler "Portal/Core/pkg/KafkaServices"
	"github.com/joho/godotenv"
)

type Config struct {
	Kafka KafkaHandler.Config
	CH    ClickhouseServices.Config
}

func GetConfig() *Config {
	err := godotenv.Load("././.env.local")

	if err != nil {
		err := godotenv.Load("././.env")

		if err != nil {
			panic(err)
		}
	}

	return &Config{
		Kafka: *KafkaHandler.GetConfig(),
		CH:    *ClickhouseServices.GetConfig(),
	}
}
