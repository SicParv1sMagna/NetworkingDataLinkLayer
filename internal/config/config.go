package config

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config Структура конфигурации;
// Содержит все конфигурационные данные о сервисе;
// автоподгружается при изменении исходного файла config.toml
type Config struct {
	ServiceHost string
	ServicePort int
	BaseURL     string
	ErrorLevel  string
}

// NewConfig Создаёт новый объект конфигурации, загружая данные из файла конфигурации
func NewConfig(ctx context.Context) (*Config, error) {
	var err error

	configName := "config"
	_ = godotenv.Load()
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}

	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.WatchConfig()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	level, err := log.ParseLevel(cfg.ErrorLevel)
	if err != nil {
		panic(err)
	}

	log.SetLevel(level)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	
	log.Info("config parsed")
	log.Println(cfg)
	
	return cfg, nil
}
