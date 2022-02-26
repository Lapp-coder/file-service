package config

import (
	"os"

	"github.com/spf13/viper"
)

const configName = "config"

type Config struct {
	Server Server `mapstructure:"server"`
	MinIO  MinIO  `mapstructure:"minio"`
}

type Server struct {
	Name         string `mapstructure:"name"`
	Host         string `mapstructure:"host"`
	Port         uint16 `mapstructure:"port"`
	BodyLimit    int    `mapstructure:"body_limit"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type MinIO struct {
	Host      string `mapstructure:"host"`
	Port      uint16 `mapstructure:"port"`
	AccessKey string
	SecretKey string
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

func New(configPath string) (Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	if err := loadEnv(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
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
