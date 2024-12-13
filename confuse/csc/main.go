package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/app-stone/confuse/csc/db"
	"github.com/HsiaoCz/app-stone/confuse/csc/handlers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("get env failed,please check the env file")
		os.Exit(1)
	}
	if err := db.Init(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("init sqlite error.....")
		os.Exit(1)
	}

	var (
		port         = os.Getenv("PORT")
		app          = http.NewServeMux()
		userHandlers = &handlers.UserHandlers{}
	)

	app.HandleFunc("POST /api/v1/user", handlers.TransferHandlerfunc(userHandlers.HandleCreateUser))

	srv := http.Server{
		Addr:         port,
		Handler:      app,
		ReadTimeout:  1500 * time.Millisecond,
		WriteTimeout: 1500 * time.Millisecond,
	}
	// 在一个goroutine中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithFields(logrus.Fields{
				"auther":       "HsiaoCz",
				"email":        "shawdml5@gmail.com",
				"phone-number": "15971078584",
				"error":        err,
			}).Error("the server is fucked up, please contact me")
			os.Exit(1)
		}
	}()
	logrus.WithFields(logrus.Fields{
		"auther":         "HsiaoCz",
		"listen address": port,
		"app name":       "book-store",
		"version":        "v0.0.1",
	}).Info("server is running .......")
	// 设置系统信号通道监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞，直到收到停止信号
	<-quit
	logrus.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 优雅关闭服务器
	if err := srv.Shutdown(ctx); err != nil {
		logrus.WithFields(logrus.Fields{
			"auther":       "HsiaoCz",
			"email":        "shawdml5@gmail.com",
			"phone-number": "15971078584",
			"error":        err,
		}).Error("the server forced to shutdown, please contact me")
		os.Exit(1)
	}

	logrus.Info("http server exit")
}
