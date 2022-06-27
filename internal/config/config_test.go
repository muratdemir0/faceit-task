package config_test

import (
	"github.com/muratdemir0/faceit-task/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {
	t.Run("Given config path and name when call new then return config", func(t *testing.T) {
		path := "../../testdata"
		name := "testconfig"

		actualConfig, _ := config.New(path, name)

		expectedConfig := &config.Config{
			Appname: "faceit-test",
			Server: config.Server{
				Port: ":3001",
			},
			Mongo: config.Mongo{
				Host: "localhost",
				Name: "faceit-task-dev",
				Port: 27017,
				Collections: config.Collections{
					Users: "users",
				},
			},
			Kafka: config.Kafka{
				URL: "kafka:9092",
			},
		}

		assert.Equal(t, expectedConfig, actualConfig)
	})

	t.Run("Given wrong config path and name when call new then return error", func(t *testing.T) {
		path := "../../wrongpath"
		name := "testconfig"
		_, err := config.New(path, name)
		assert.NotNil(t, err)
	})
}
