package config

import "github.com/ilyakaznacheev/cleanenv"

var (
	config *Config
)

type KafkaInstance struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
	GroupID string   `yaml:"groupId"`
}
type InfluxInstance struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Token    string `yaml:"token"`
}
type Config struct {
	KafkaInstance  KafkaInstance  `yaml:"kafkaInstance"`
	InfluxInstance InfluxInstance `yaml:"influxInstance"`
}

func LoadConfiguration() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	config = &cfg
	return &cfg, nil
}

func GetConfig() *Config {
	return config
}
