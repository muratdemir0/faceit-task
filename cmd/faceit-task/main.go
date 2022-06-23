package main

import (
	"context"
	"fmt"
	"github.com/muratdemir0/faceit-task/internal/config"
	"github.com/muratdemir0/faceit-task/internal/user"
	"github.com/muratdemir0/faceit-task/pkg/server"
	"github.com/muratdemir0/faceit-task/pkg/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	appEnv := os.Getenv("APP_ENV")
	conf, err := config.New(".config", appEnv)
	if err != nil {
		return err
	}
	conf.Print()

	logger, err := zap.NewProduction()
	defer func() {
		err = logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}()

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Mongo.User, conf.Mongo.Password,
		conf.Mongo.Host, conf.Mongo.Port)
	mongoClient, connectErr := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if connectErr != nil {
		return connectErr
	}
	if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Successfully connected and pinged.")

	userStore := store.NewUserStore(mongoClient, &conf.Mongo)
	userService := user.NewService(userStore)
	userHandler := user.NewHandler(userService)

	handlers := []server.Handler{userHandler}

	s := server.New(conf.Server.Port, handlers, logger)

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	go s.Run()

	<-shutdownChan

	s.Stop()
	return nil
}
