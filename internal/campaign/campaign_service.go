package campaign

import (
	"context"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type service struct {
	timeout time.Duration
	DB      *gorm.DB
}

func NewService() Service {
	return &service{
		time.Duration(20) * time.Second,
		config.DB,
	}
}

func (s *service) CreateCampaign(c context.Context, req *CreateCampaignReq) (*CreateCampaignRes, error) {
	var campaign = models.Campaign{
		UserID:             req.UserID,
		Name:               req.Name,
		Description:        req.Description,
		RunEveryNMinutes:   req.RunEveryNMinutes,
		TotalProspects:     req.TotalProspects,
		ProcessedProspects: 0,
		ProcessedAt:        nil,
		CurrentStatus:      "CREATING",
	}
	s.DB.Create(&campaign)
	logger.Logger.Info("Campaign created", zap.Any("Campaign", campaign))

	return &CreateCampaignRes{
		ID:                 campaign.ID,
		UserID:             campaign.UserID,
		Name:               campaign.Name,
		Description:        campaign.Description,
		RunEveryNMinutes:   campaign.RunEveryNMinutes,
		TotalProspects:     campaign.TotalProspects,
		ProcessedProspects: campaign.ProcessedProspects,
		ProcessedAt:        campaign.ProcessedAt,
		CurrentStatus:      campaign.CurrentStatus,
	}, nil
}
