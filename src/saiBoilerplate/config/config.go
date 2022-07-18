package config

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	HttpServer struct {
		Host string
		Port string
	}
	Address struct {
		Url string
	}
	SocketServer  struct {
		Host string
		Port string
	}
	Token   string
	Storage struct {
		Atlas    bool
		User     string
		Pass     string
		Host     string
		Port     string
		Database string
	}
	Operations []string
	StartBlock int
	WebSocket  struct {
		Token string
		Url   string
	}
	Contract struct {
		Address string
		ABI     string
	}
	Geth  []string
	Sleep int
}

func Load() Configuration {
	var config Configuration
	err := gonfig.GetConf("config.json", &config)

	if err != nil {
		fmt.Println("Configuration problem:", err)
		panic(err)
	}

	return config
}
