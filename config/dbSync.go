package config

import (
	"log"

	"github.com/anuragdaksh7/zapmail-backend/models"
)

func SyncDB() {
	err := DB.AutoMigrate(
		models.User{},
		models.UserToken{},
		models.Campaign{},
		models.Prospect{},
	)
	if err != nil {
		log.Fatal(err)
	}
}
