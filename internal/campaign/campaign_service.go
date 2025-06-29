package campaign

import (
	"context"
	"errors"
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
		Description:        "",
		RunEveryNMinutes:   20,
		TotalProspects:     0,
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

func (s *service) GetCampaigns(c context.Context, req *GetCampaignsReq) (*GetCampaignsRes, error) {
	var campaigns []models.Campaign
	var userCampaigns []Campaign
	s.DB.
		Where("user_id = ?", req.UserID).
		Find(&campaigns)
	logger.Logger.Info("Num Campaigns ", zap.Int("Campaigns", len(campaigns)))

	for _, campaign := range campaigns {
		userCampaigns = append(userCampaigns, Campaign{
			ID:                 campaign.ID,
			UserID:             campaign.UserID,
			Name:               campaign.Name,
			Description:        campaign.Description,
			RunEveryNMinutes:   campaign.RunEveryNMinutes,
			TotalProspects:     campaign.TotalProspects,
			ProcessedProspects: campaign.ProcessedProspects,
			ProcessedAt:        campaign.ProcessedAt,
			CurrentStatus:      campaign.CurrentStatus,
			CreatedAt:          campaign.CreatedAt,
			UpdatedAt:          campaign.UpdatedAt,
		})
	}

	return &GetCampaignsRes{userCampaigns}, nil
}

func (s *service) CreateCampaignWithProspects(c context.Context, req *CreateCampaignWithProspectsReq) (*CreateCampaignRes, error) {
	var newCampaign = models.Campaign{
		UserID:             req.UserID,
		Name:               req.Name,
		Description:        "",
		RunEveryNMinutes:   20,
		TotalProspects:     uint(len(req.Prospects)),
		ProcessedProspects: 0,
		ProcessedAt:        nil,
		CurrentStatus:      "RUNNING",
	}
	s.DB.Save(&newCampaign)
	logger.Logger.Info("Campaign created", zap.Any("Campaign", newCampaign))

	go func() {
		var prospects []models.Prospect
		logger.Logger.Info("prospects found: ", zap.Int("prospects", len(prospects)))
		for _, prospect := range req.Prospects {
			prospects = append(prospects, models.Prospect{
				CampaignID:  newCampaign.ID,
				FullName:    prospect.FullName,
				Email:       prospect.Email,
				Phone:       prospect.Phone,
				Company:     prospect.Company,
				Position:    prospect.Position,
				LinkedInURL: prospect.LinkedInURL,
				Status:      prospect.Status,
				Location:    prospect.Location,
				LastContact: prospect.LastContact,
				ResponseAt:  prospect.ResponseAt,
				Notes:       prospect.Notes,
			})
		}
		logger.Logger.Info("prospects ready to create: ", zap.Int("prospects", len(prospects)))
		s.DB.Create(&prospects)
	}()

	return &CreateCampaignRes{
		ID:                 newCampaign.ID,
		UserID:             newCampaign.UserID,
		Name:               newCampaign.Name,
		Description:        newCampaign.Description,
		RunEveryNMinutes:   newCampaign.RunEveryNMinutes,
		TotalProspects:     newCampaign.TotalProspects,
		ProcessedProspects: newCampaign.ProcessedProspects,
		ProcessedAt:        newCampaign.ProcessedAt,
		CurrentStatus:      newCampaign.CurrentStatus,
	}, nil
}

func (s *service) ToggleTemplateAssociation(c context.Context, req *ToggleTemplateAssociationReq) (*ToggleTemplateAssociationRes, error) {
	var templateCampaign TemplateCampaign
	var template models.Template
	var campaign models.Campaign

	if err := s.DB.First(&template, req.TemplateID).Error; err != nil {
		return nil, err // or handle record not found
	}

	// Load campaign
	if err := s.DB.First(&campaign, req.CampaignID).Error; err != nil {
		return nil, err
	}

	// Check ownership
	if template.UserID != req.UserID || campaign.UserID != req.UserID {
		logger.Logger.Info("template campaign user id does not match template campaign user id")
		return nil, errors.New("unauthorized access to template or campaign")
	}

	err := s.DB.Where("template_id = ? AND campaign_id = ?", req.TemplateID, req.CampaignID).First(&templateCampaign).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Logger.Info("template campaign not found")
		s.DB.Table("template_campaigns").Create(&TemplateCampaign{
			TemplateID: req.TemplateID,
			CampaignID: req.CampaignID,
		})
		logger.Logger.Info("created new template campaign association")
	} else if err != nil {
		logger.Logger.Error("Error on ToggleTemplateAssociation", zap.Error(err))
	} else {
		s.DB.Unscoped().Table("template_campaigns").
			Where("template_id = ? AND campaign_id = ?", req.TemplateID, req.CampaignID).
			Delete(&TemplateCampaign{})
	}

	s.DB.
		Preload("Template").
		Where("id = ?", req.CampaignID).
		First(&campaign)

	return &ToggleTemplateAssociationRes{}, nil
}
