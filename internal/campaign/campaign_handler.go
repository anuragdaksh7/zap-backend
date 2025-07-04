package campaign

import (
	"fmt"
	"net/http"

	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) CreateCampaign(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		logger.Logger.Warn("Failed to get user from context : UNAUTHORIZED")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)
	var req CreateCampaignReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("Failed to bind request: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	req.UserID = currentUser.ID
	fmt.Println(req)

	res, err := h.Service.CreateCampaign(c, &req)
	if err != nil {
		logger.Logger.Error("Failed to create campaign: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) GetCampaigns(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		logger.Logger.Warn("Failed to get user from context : UNAUTHORIZED")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)

	var req GetCampaignsReq
	req.UserID = currentUser.ID

	res, err := h.Service.GetCampaigns(c, &req)
	if err != nil {
		logger.Logger.Error("Failed to get campaigns: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) CreateCampaignWithProspects(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		logger.Logger.Warn("Failed to get user from context : UNAUTHORIZED")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)

	var req CreateCampaignWithProspectsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("Failed to bind request: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = currentUser.ID

	res, err := h.Service.CreateCampaignWithProspects(c, &req)
	if err != nil {
		logger.Logger.Error("Failed to create campaigns: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func (h *Handler) ToggleCampaignTemplateAssociations(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		logger.Logger.Warn("Failed to get user from context : UNAUTHORIZED")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)

	var req ToggleTemplateAssociationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Error("Failed to bind request: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserID = currentUser.ID

	res, err := h.Service.ToggleTemplateAssociation(c, &req)
	if err != nil {
		logger.Logger.Error("Failed to toggle campaign template associations: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
