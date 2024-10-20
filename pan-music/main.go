package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/app-stone/pan-music/handlers"
	"github.com/HsiaoCz/app-stone/pan-music/handlers/middlewares"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("get env failed,please check the env file")
		os.Exit(1)
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	var (
		port        = os.Getenv("PORT")
		testHandler = &handlers.TestHandler{}
		userHandler = &handlers.UserHandlers{}
		router      = http.NewServeMux()
	)
	router.HandleFunc("GET /api/v1/test", testHandler.HandleTestConnect)
	router.HandleFunc("POST /api/v1/user", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(userHandler.HandleCreateUser)))

	srv := http.Server{
		Addr:         port,
		Handler:      router,
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
		"app name":       "pan-music",
	}).Info("server started .......")
	// 设置系统信号通道监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞，直到收到停止信号
	<-quit
	logrus.Info("Shutting down server...")

	// 创建一个上下文，设置超时时间，确保优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
