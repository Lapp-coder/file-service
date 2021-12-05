package config

import (
	"os"

	"github.com/spf13/viper"
)

const configName = "config"

type Config struct {
	Server
	MinIO
}

type Server struct {
	Name         string `mapstructure:"name"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	BodyLimit    int    `mapstructure:"body_limit"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type MinIO struct {
	AccessKey string
	SecretKey string
}

func New(configPath string) (Config, error) {
	var cfg Config
	if err := unmarshal(configPath, &cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func unmarshal(configPath string, cfg *Config) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return err
	}

	if err := loadEnv(cfg); err != nil {
		return err
	}

	return nil
}

func loadEnv(cfg *Config) error {
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	if accessKey == "" {
		return ErrMinIOAccessKeyIsEmpty
	}

	secretKey := os.Getenv("MINIO_SECRET_KEY")
	if secretKey == "" {
		return ErrMinIOSecretKeyIsEmpty
	}

	cfg.MinIO.AccessKey = accessKey
	cfg.MinIO.SecretKey = secretKey

	return nil
}
