package config

import (
	"../system"
	"encoding/json"
	"log"
	"os"
)

const (
	DefaultHost         string = "127.0.0.1"
	DefaultPort         uint16 = 80
	DefaultService      string = "Default Service: tcp proxy"
	DefaultStrategy     string = "iphash"
	DefaultMaxProcessor int    = 2
	DefaultHeartbeat    int    = 50000
	DefaultLogFile      string = "/tmp/logs"
	DefaultProtocal     string = "tcp"
)

// description
// Config
type Config struct {
	Protocal         string           `json:"protocal"`
	Host             string           `json:"host"`
	Port             uint16           `json:"port"`
	Service          string           `json:"service"`
	Strategy         string           `json:"strategy"`
	RequestQueueSize int              `json:"requestqueuesize"`
	MaxProcessor     int              `json:"maxprocessor"`
	Heartbeat        int              `json:"heartbeat"`
	Keepalive        int              `json:"keepalive"`
	LogFile          string           `json:"logfile"`
	Backends         []system.Backend `json:"backends"`
}

// description
// load config file
func Load(filename string) (*Config, error) {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		log.Println("load config error:", err)
	} else {
		buff := make([]byte, 1024)
		end, _ := file.Read(buff)
		err = json.Unmarshal(buff[:end], &config)
		if err != nil {
			log.Println("decode json error:", err)
		}
	}

	// add default value for config file
	if config.Host == "" {
		config.Host = DefaultHost
		log.Println("success load default config parameter Host", config.Host)
	}
	if config.Port == 0 {
		config.Port = DefaultPort
		log.Println("success load default config parameter Port", config.Port)
	}
	if config.Heartbeat == 0 {
		config.Heartbeat = DefaultHeartbeat
		log.Println("success load default config parameter Heartbeat", config.Heartbeat)
	}
	if config.Strategy == "" {
		config.Strategy = DefaultStrategy
		log.Println("success load default config parameter Strategy", config.Strategy)
	}
	if config.MaxProcessor == 0 {
		config.MaxProcessor = DefaultMaxProcessor
		log.Println("success load default config parameter MaxProcessor", config.MaxProcessor)
	}
	if config.LogFile == "" {
		config.LogFile = DefaultLogFile
		log.Println("success load default config parameter LogFile", config.LogFile)
	}
	if config.DefaultProtocal == "" {
		config.DefaultProtocal = DefaultProtocal
		log.Println("success load default config parameter Protocal", config.DefaultProtocal)
	}

	log.Println("success load config file:", filename)
	return &config, err
}
