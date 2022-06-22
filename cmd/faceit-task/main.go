package main

import (
	"fmt"
	"github.com/muratdemir0/faceit-task/internal/config"
	"os"
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
	return nil
}
