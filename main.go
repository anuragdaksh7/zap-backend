package main

import (
	"log"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/router"
)

var _config config.Config

func init() {
	var err error
	_config, err = config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	config.ConnectDB()
	config.SyncDB()
	config.InitRedis()
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
}

func main() {
	router.InitRouter()
	log.Fatal(router.Start("0.0.0.0:" + _config.PORT))
}
