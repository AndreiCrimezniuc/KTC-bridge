package ClickhouseServices

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	RetryTimeout time.Duration `yaml:"retryTimeout"`
	Database     string        `yaml:"database"`
}

func Connection(config *Config) *ch.DB {
	db := ch.Connect(
		ch.WithAddr(os.Getenv("CH_SERVER_ADDRESS")),
		ch.WithInsecure(true),
		ch.WithDatabase(config.Database),
	)
	db.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))

	return db
}

func Write(data []any, ctx context.Context, db *ch.DB, config Config) error {
	var result sql.Result
	var err error

	for {
		result, err = db.NewInsert().Model(&data).Exec(ctx)

		if err == nil {
			break
		} else {
			log.Println("Error in writing to Clickhouse. Trying again... ", err)
			time.Sleep(config.RetryTimeout)
		}
	}

	rowsAffected, _ := result.RowsAffected()

	fmt.Printf("Inserted : %d \n", rowsAffected)

	return nil
}

func Ping(c *Config, db *ch.DB, ctx context.Context) {
	for {
		err := db.Ping(ctx)

		if err == nil {
			break
		} else {
			log.Println("The error occurred while pinging to clickhouse. Trying again... ", err)
			time.Sleep(c.RetryTimeout)
		}
	}
}

func GetConfig() *Config {
	fh, err := os.Open("././configs/clickhouse_config.yaml")

	if err != nil {
		log.Fatalln("can't open Clickhouse config in 'configs/clickhouse_config.yaml'")
	}

	chConf := Config{}

	if err = yaml.NewDecoder(fh).Decode(&chConf); err != nil {
		log.Fatalln(err)
	}

	return &chConf
}
