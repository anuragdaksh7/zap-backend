package template

import (
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

func (h *Handler) CreateTemplate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		logger.Logger.Warn("Failed to get user from context : UNAUTHORIZED")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	currentUser := user.(models.User)

	var req CreateTemplateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Logger.Warn("Failed to bind request", zap.Any("request", req))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	req.UserID = currentUser.ID

	res, err := h.Service.CreateTemplate(c, &req)
	if err != nil {
		logger.Logger.Warn("Failed to create template")
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
