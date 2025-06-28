package jobs

import (
	"fmt"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/models"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func StartCronJobs() {
	c := cron.New(cron.WithSeconds()) // allow second-level scheduling

	// Example: Runs every 10 seconds
	c.AddFunc("*/10 * * * * *", func() {
		fmt.Println("🔁 Scheduled task at:", time.Now().Format(time.RFC3339))
		// Do something: update DB, send mail, etc.
	})

	c.AddFunc("0 */10 * * * *", func() {
		logger.Logger.Info("SENDING MAILs: ")
		var campaigns []models.Campaign
		config.DB.Where("current_status = ?", "RUNNING").Find(&campaigns)
		logger.Logger.Info("campaigns found: ", zap.Int("campaigns", len(campaigns)))
	})

	c.Start()
}
