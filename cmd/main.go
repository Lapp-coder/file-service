package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	fileHandler "github.com/Lapp-coder/file-service/internal/adapters/api/v1/file"

	fileService "github.com/Lapp-coder/file-service/internal/domain/file"

	fileStorage "github.com/Lapp-coder/file-service/internal/adapters/db/file"

	"github.com/Lapp-coder/file-service/internal/config"
	"github.com/Lapp-coder/file-service/pkg/client"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const configPath = "configs/"

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.SetReportCaller(true)

	cfg, err := config.New(configPath)
	if err != nil {
		logrus.Fatalf("failed to init config: %s", err.Error())
	}

	app := fiber.New(fiber.Config{
		AppName:      cfg.Server.Name,
		BodyLimit:    cfg.Server.BodyLimit << 20, // MB
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	})

	minioClient, err := client.NewMinIOClient(cfg.MinIO)
	if err != nil {
		logrus.Fatalf("failed to init minio client: %s", err.Error())
	}

	postgresConn, err := client.NewPostgresConn(cfg.Postgres)
	if err != nil {
		logrus.Fatalf("failed to init connection with postgres: %s", err.Error())
	}

	storage := fileStorage.New(minioClient, postgresConn)
	service := fileService.New(storage)
	handler := fileHandler.New(service)

	handler.Register(app.Group("/api/v1"))

	go func() {
		addr := cfg.Server.Host + ":" + strconv.Itoa(int(cfg.Server.Port))
		if err = app.Listen(addr); err != nil {
			logrus.Errorf("failed to start server: %s", err.Error())
		}
	}()

	logrus.Info("file-service started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("file-service shutdown")

	if err = app.Shutdown(); err != nil {
		logrus.Errorf("failed to gracefully shutdown file-service: %s", err.Error())
	}

	if err = postgresConn.Close(); err != nil {
		logrus.Errorf("failed to close connection with postgres: %s", err.Error())
	}
}
