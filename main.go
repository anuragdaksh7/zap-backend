package main

import (
	"log"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/internal/ai"
	"github.com/anuragdaksh7/zapmail-backend/internal/campaign"
	"github.com/anuragdaksh7/zapmail-backend/internal/oAuth"
	"github.com/anuragdaksh7/zapmail-backend/internal/template"
	"github.com/anuragdaksh7/zapmail-backend/jobs"
	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/router"
	"github.com/anuragdaksh7/zapmail-backend/utils"
)

var _config config.Config

func init() {
	var err error
	_config, err = config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	utils.InitEncryptionKey(_config.EncryptionKey)

	logger.InitLogger(_config)
	logger.Logger.Info("Logger initialized")
	config.ConnectDB()
	logger.Logger.Info("DB connection established")
	config.SyncDB()
	logger.Logger.Info("DB sync completed")
	config.InitRedis()
	logger.Logger.Info("Redis connected")
	ai.InitGemini()
	logger.Logger.Info("Gemini connected")
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

	jobs.StartCronJobs()

	oAuthSvc := oAuth.NewService()
	campaignSvc := campaign.NewService()
	templateSvc := template.NewService()

	oAuthHandler := oAuth.NewHandler(oAuthSvc)
	campaignHandler := campaign.NewHandler(campaignSvc)
	templateHandler := template.NewHandler(templateSvc)

	router.InitRouter(oAuthHandler, campaignHandler, templateHandler)
	log.Fatal(router.Start("0.0.0.0:" + _config.PORT))
}
