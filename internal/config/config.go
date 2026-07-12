package config

import (
	"flag"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServerConfig struct {
	Port string `yaml:"port" env:"PORT" env-default:":7505" `
	Host string `yaml:"host" env:"HOST" env-default:"localhost"`
}

type DBConfig struct {
	Port uint16 `yaml:"port" env:"DB_PORT"`
	User string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Database string `yaml:"database" env:"DATABASE"`
} 
type Config struct {
	HttpServerConfig `yaml:"http_server" env:"HTTP_SERVER"`
	DBConfig `yaml:"db" env:"DBCONFIG"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configFlag := flag.String("config","","path to config file")
		flag.Parse()

		configPath = *configFlag

		if configPath == "" {
			slog.Error("config must load: configPath not set")
			os.Exit(1)
		}
	}

	_,err := os.Lstat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Error("config must load: configPath does not exist","error",err)
			os.Exit(1)
		}
	}

	var config Config

	err = cleanenv.ReadConfig(configPath, &config)
	if err != nil {
		slog.Error("config must load: failed to read config", "error",err)
		os.Exit(1)
	}
	return &config
	
}
