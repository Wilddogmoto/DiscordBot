package cfg

import (
	"encoding/json"
	"github.com/Wilddogmoto/DiscordBot/loggds"
	"os"
)

type Config struct {
	DataBase struct {
		Server   string `json:"server"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"DB"`
	}
	Discord struct {
		Token string `json:"token"`
	}
}

var Parameter *Config

func Configuration(filename string) {

	var (
		log = loggds.Logg
		err error
		cfg *os.File
		res []byte
	)

	// open file config and decode
	if cfg, err = os.Open(filename); err != nil {
		log.Fatalf("error open cfg file: %s", err)
	}

	if err = json.NewDecoder(cfg).Decode(&Parameter); err != nil {
		log.Fatalf("error on decode config file: %v", err)
	}

	if res, err = json.Marshal(&Parameter); err != nil {
		log.Errorf("error on encoded json: %v", err)
	}

	log.Infof("Start discord bot with parameters: %s", res)
}
