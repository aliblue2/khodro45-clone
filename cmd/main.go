package main

import (
	"github.com/aliblue2/khodro45/api"
	"github.com/aliblue2/khodro45/configs"
	"github.com/aliblue2/khodro45/data/cache"
	"github.com/aliblue2/khodro45/data/db"
	"github.com/aliblue2/khodro45/pkg/logging"
)

func main() {
	cfg := configs.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitiateRedis(cfg)
	defer cache.CloseRedis()

	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = db.InitiateDb(cfg)
	defer db.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
		panic(err)
	}

	api.InitiateServer(cfg)
}
