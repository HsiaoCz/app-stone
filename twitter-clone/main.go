package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/app-stone/twitter-clone/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Error("Error loading .env file")
		os.Exit(1)
	}

	if err := db.InitDB(); err != nil {
		logrus.WithError(err).Error("Error initializing database")
		os.Exit(1)
	}

	var (
		port = os.Getenv("PORT")
	)

	// Start the server
	app := fiber.New()

	// Routes

	// Listen on port 3000
	// 创建一个信号通道
	quit := make(chan os.Signal, 1)

	// 监听 SIGINT 和 SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 启动 Fiber 应用
	go func() {
		if err := app.Listen(port); err != nil {
			logrus.WithError(err).Error("Error starting server")
			os.Exit(1)
		}
	}()

	logrus.Infof("Server running on port %s", port)
	// 等待信号
	<-quit
	logrus.Info("Shutting down server")
	// 创建上下文用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭 Fiber 应用
	if err := app.ShutdownWithContext(ctx); err != nil {
		logrus.WithError(err).Error("Error shutting down server")
		os.Exit(1)
	}

	logrus.Info("Server shutdown successfully")
}
