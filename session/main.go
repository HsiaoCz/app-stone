package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/HsiaoCz/app-stone/session/db"
	"github.com/HsiaoCz/app-stone/session/handlers"
	"github.com/HsiaoCz/app-stone/session/storage"
	"github.com/gofiber/fiber/v2"
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

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if e, ok := err.(handlers.ErrorMsg); ok {
			return c.Status(e.Status).JSON(&e)
		}
		aErr := handlers.ErrorMsg{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
		return c.Status(aErr.Status).JSON(&aErr)
	},
}

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": err.Error(),
		}).Error("get env failed,please check it out.....")
		os.Exit(1)
	}

	if err := db.Init(); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": err.Error(),
		}).Error("db init error,please check it out.....")
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGOURL")))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"message": err.Error(),
		}).Error("connect mongodb error,please check it out.....")
		os.Exit(1)
	}

	go func() {
		if err := client.Ping(ctx, &readpref.ReadPref{}); err != nil {
			logrus.WithFields(logrus.Fields{
				"message": err.Error(),
			}).Error("ping mongodb error,please check it out.....")
			os.Exit(1)
		}
	}()

	var (
		app     = fiber.New(config)
		port    = os.Getenv("PORT")
		userApp = handlers.UserHandlersInit(storage.UserStoreInit(client, client.Database(os.Getenv("DBNAME")).Collection(os.Getenv("USERCOLL"))), storage.SessionStoreInit(db.Get()))
		v1      = app.Group("/api/v1")
	)

	{
		// user handler function
		v1.Post("/user", userApp.HandleCreateUser)
	}

	// restart and shutdown
	go func() {
		if err := app.Listen(port); err != nil {
			logrus.WithFields(logrus.Fields{
				"message": err.Error(),
			}).Error("server listen and serve error,please check out ......")
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := app.Shutdown(); err != nil {
		logrus.WithFields(logrus.Fields{
			"message": err.Error(),
		}).Error("server shutdown error,please check out ......")
		os.Exit(1)
	}
}
