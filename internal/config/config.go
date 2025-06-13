package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true" `
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address string
}

func MustLoad() *Config {
	var configPath string
	// get path from env file or
	configPath = os.Getenv("CONFIG_PATH")

	// get path from flag
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path not found !")
		}
	}

	// check whether the file exist in path or not
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Fatalf("Config file not found : %s", configPath)
	}

	// load configuration and return
	var cfg Config
	err = cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
