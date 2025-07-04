package router

import (
	"net/http"

	"github.com/anuragdaksh7/zapmail-backend/internal/campaign"
	"github.com/anuragdaksh7/zapmail-backend/internal/email"
	"github.com/anuragdaksh7/zapmail-backend/internal/oAuth"
	"github.com/anuragdaksh7/zapmail-backend/internal/template"
	"github.com/anuragdaksh7/zapmail-backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(
	oAuthHandler *oAuth.Handler,
	campaignHandler *campaign.Handler,
	templateHandler *template.Handler,
	emailHandler *email.Handler,
) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://resourcify-three.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	oAuthRouter := r.Group("/oauth")
	oAuthRouter.GET("google/redirect-uri", oAuthHandler.GetRedirectURL)
	oAuthRouter.GET("/google/callback", oAuthHandler.HandleGoogleCallback)

	campaignRouter := r.Group("/campaign")
	campaignRouter.POST("", middleware.RequireAuth, campaignHandler.CreateCampaign)
	campaignRouter.GET("", middleware.RequireAuth, campaignHandler.GetCampaigns)
	campaignRouter.POST("/prospects", middleware.RequireAuth, campaignHandler.CreateCampaignWithProspects)
	campaignRouter.POST("/template", middleware.RequireAuth, campaignHandler.ToggleCampaignTemplateAssociations)

	templateRouter := r.Group("/template")
	templateRouter.POST("", middleware.RequireAuth, templateHandler.CreateTemplate)

	emailRouter := r.Group("/email")
	emailRouter.GET("", middleware.RequireAuth, emailHandler.SyncEmails)
}

func Start(addr string) error {
	return r.Run(addr)
}
