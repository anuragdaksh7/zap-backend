package main

import (
	"log"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/router"
)

var _config config.Config

func init() {
	var err error
	_config, err = config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	logger.InitLogger(_config)
	logger.Logger.Info("Logger initialized")
	config.ConnectDB()
	logger.Logger.Info("DB connection established")
	config.SyncDB()
	logger.Logger.Info("DB sync completed")
	config.InitRedis()
	logger.Logger.Info("Redis connected")
	defer logger.Logger.Sync()
}

func main() {
	defer func() {
		if config.RedisClient != nil {
			err := config.RedisClient.Close()
			if err != nil {
				panic(err)
				return
			}
			log.Println("Redis client closed.")
		}
	}()
	router.InitRouter()
	log.Fatal(router.Start("0.0.0.0:" + _config.PORT))
}
