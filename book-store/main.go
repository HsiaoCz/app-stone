package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/HsiaoCz/app-stone/book-store/db"
	"github.com/HsiaoCz/app-stone/book-store/handlers"
	"github.com/HsiaoCz/app-stone/book-store/handlers/middlewares"
	"github.com/HsiaoCz/app-stone/book-store/storage"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	// init sqlite

	if err := db.Init(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("init sqlite error.....")
		os.Exit(1)
	}

	// init redis

	go func() {
		count, err := strconv.Atoi(os.Getenv("DBCOUNT"))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"error message": err,
			}).Error("dbcount must be int")
			os.Exit(1)
		}
		db.InitRedis(os.Getenv("REDISURL"), os.Getenv("PASSWD"), count)
	}()

	// 创建一个上下文，设置超时时间，确保优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error message": err,
		}).Error("connect mongodb failed,please check the env file....")
		os.Exit(1)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			logrus.WithFields(logrus.Fields{
				"error message": err,
			}).Error("mongodb connection broken,please check out....")
			os.Exit(1)
		}
	}()

	var (
		port          = os.Getenv("PORT")
		testHandler   = &handlers.TestHandler{}
		userHandler   = handlers.UserHandlersInit(storage.UserDataInit(db.Get()))
		bookHandler   = handlers.BookHandlersInit(storage.BookDataInit(db.Get()))
		reviewHandler = handlers.ReviewHandlersInit(storage.ReviewDataInit(db.Get()))
		router        = http.NewServeMux()
	)

	// router
	router.HandleFunc("GET /api/v1/test", testHandler.HandleTestConnect)
	router.HandleFunc("POST /api/v1/user", handlers.TransferHandlerfunc(userHandler.HandleCreateUser))
	router.HandleFunc("POST /api/v1/user/login", handlers.TransferHandlerfunc(userHandler.HandleUserLogin))
	router.HandleFunc("GET /api/v1/user/{user_id}", handlers.TransferHandlerfunc(userHandler.HandleGetUserByID))
	router.HandleFunc("DELETE /api/v1/user/{user_id}", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(userHandler.HandleDeleteUserByID)))
	router.HandleFunc("PUT /api/v1/user", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(userHandler.HandleUpdateUser)))
	router.HandleFunc("GET /api/v1/user/search", handlers.TransferHandlerfunc(userHandler.HandleGetUserByUsername))

	router.HandleFunc("POST /api/v1/book", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(bookHandler.HandleCreateBook)))
	router.HandleFunc("GET /api/v1/book/{book_id}", handlers.TransferHandlerfunc(bookHandler.HandleGetBookByID))
	router.HandleFunc("DELETE /api/v1/book/{book_id}", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(bookHandler.HandleDeleteBookByID)))
	router.HandleFunc("PUT /api/v1/book", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(bookHandler.HandleUpdateBook)))

	router.HandleFunc("POST /api/v1/review", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(reviewHandler.HandleCreateReview)))
	router.HandleFunc("GET /api/v1/review/{book_id}", middlewares.JwtMiddleware(handlers.TransferHandlerfunc(reviewHandler.HandleGetReviewByBookID)))

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
		"app name":       "book-store",
		"version":        "v0.0.1",
	}).Info("server is running .......")
	// 设置系统信号通道监听
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞，直到收到停止信号
	<-quit
	logrus.Info("Shutting down server...")

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
