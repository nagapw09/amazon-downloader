package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Host         string
	Port         int
	WorkPoolSize int
	MaxQueueSize int
}

func NewConfig() (*Config, error) {
	host, ok := os.LookupEnv("AMAZON_DOWNLOADER_HOST")
	if !ok {
		return nil, fmt.Errorf("AMAZON_DOWNLOADER_HOST is empty")
	}

	rawPort, ok := os.LookupEnv("AMAZON_DOWNLOADER_PORT")
	if !ok {
		return nil, fmt.Errorf("AMAZON_DOWNLOADER_PORT is empty")
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		return nil, fmt.Errorf("AMAZON_DOWNLOADER_PORT is not a string")
	}

	var poolSize = 10
	rawPoolSize, ok := os.LookupEnv("AMAZON_DOWNLOADER_POOL_SIZE")
	if ok {
		poolSize, err = strconv.Atoi(rawPoolSize)
		if err != nil {
			poolSize = 10
		}
	}

	var maxQueueSize = 10
	rawMaxQueueSize, ok := os.LookupEnv("AMAZON_DOWNLOADER_QUEUE_SIZE")
	if ok {
		poolSize, err = strconv.Atoi(rawMaxQueueSize)
		if err != nil {
			maxQueueSize = 10
		}
	}

	return &Config{
		Host:         host,
		Port:         port,
		WorkPoolSize: poolSize,
		MaxQueueSize: maxQueueSize,
	}, nil
}
