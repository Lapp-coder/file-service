package config

import (
	"os"

	"github.com/spf13/viper"
)

const configName = "config"

type Config struct {
	Server
	MinIO
	Postgres
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

	if err := viper.UnmarshalKey("minio", &cfg.MinIO); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
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

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return ErrPostgresPasswordIsEmpty
	}

	cfg.MinIO.AccessKey = accessKey
	cfg.MinIO.SecretKey = secretKey
	cfg.Postgres.Password = postgresPassword

	return nil
}
