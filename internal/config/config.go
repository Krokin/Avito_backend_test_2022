package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MySQL string `yaml:"MySQL_dsn"`
	Address string `yaml:"address"`
	Port string `yaml:"port"`
}

func GetConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("./build/configs/config.yml", &cfg)
	return &cfg, err
}
