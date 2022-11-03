package tests

import (
	"Portal/Core/pkg/ClickhouseServices"
	KafkaConsumer "Portal/Core/pkg/KafkaServices"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func TestKafkaConfig(t *testing.T) {
	fh, err := os.Open("./../../configs/kafka_config.yaml")
	if err != nil {
		t.Error("can't open kafka config, ", err)
	}

	kafkaConfig := KafkaConsumer.Config{}

	if er := yaml.NewDecoder(fh).Decode(&kafkaConfig); er != nil {
		t.Error("can't decode/map kafka config on struct, ", er)
	}
}

func TestCHConfig(t *testing.T) {
	fh, err := os.Open("./../../configs/clickhouse_config.yaml")

	if err != nil {
		t.Error("can't open Clickhouse config, ", err)
	}

	clickhouseConfig := ClickhouseServices.Config{}

	if er := yaml.NewDecoder(fh).Decode(&clickhouseConfig); er != nil {
		t.Error("can't decode/map clickhouse config on struct", er)
	}
}
