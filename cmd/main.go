package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Lapp-coder/file-service/internal/composites"

	"github.com/Lapp-coder/file-service/internal/client"

	"github.com/Lapp-coder/file-service/internal/config"
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

	fileComposite := composites.NewFileComposite(minioClient)

	apiV1 := app.Group("/api/v1")
	fileComposite.Handler.Init(apiV1)

	go func() {
		addr := cfg.Server.Host + ":" + strconv.Itoa(int(cfg.Server.Port))
		if err = app.Listen(addr); err != nil {
			logrus.Errorf("failed to start server: %s", err.Error())
		}
	}()

	logrus.Info("file-service started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("file-service shutdown")

	if err = app.Shutdown(); err != nil {
		logrus.Errorf("failed to gracefully shutdown file-service: %s", err.Error())
	}
}
