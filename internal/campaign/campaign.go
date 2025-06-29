package campaign

import (
	"context"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/internal/prospect"
	"github.com/anuragdaksh7/zapmail-backend/models"
)

type Campaign struct {
	ID                 uint       `json:"id"`
	UserID             uint       `json:"user_id"`
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	RunEveryNMinutes   uint       `json:"run_every_n_minutes"`
	TotalProspects     uint       `json:"total_prospects"`
	ProcessedProspects uint       `json:"processed_prospects"`
	ProcessedAt        *time.Time `json:"processed_at"`
	CurrentStatus      string     `json:"current_status"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

type CreateCampaignReq struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}

type CreateCampaignRes struct {
	ID                 uint       `json:"id"`
	UserID             uint       `json:"user_id"`
	Name               string     `json:"name"`
	Description        string     `json:"description"`
	RunEveryNMinutes   uint       `json:"run_every_n_minutes"`
	TotalProspects     uint       `json:"total_prospects"`
	ProcessedProspects uint       `json:"processed_prospects"`
	ProcessedAt        *time.Time `json:"processed_at"`
	CurrentStatus      string     `json:"current_status"`
}

type GetCampaignsReq struct {
	UserID uint `json:"user_id"`
}

type GetCampaignsRes struct {
	Campaigns []Campaign `json:"campaigns"`
}

type CreateCampaignWithProspectsReq struct {
	UserID    uint                   `json:"user_id"`
	Name      string                 `json:"name"`
	Prospects []prospect.NewProspect `json:"prospects"`
}

type TemplateCampaign struct {
	TemplateID uint
	CampaignID uint
}

type ToggleTemplateAssociationReq struct {
	UserID     uint `json:"user_id"`
	TemplateID uint `json:"template_id"`
	CampaignID uint `json:"campaign_id"`
}

type ToggleTemplateAssociationRes struct {
	Templates []models.Template `json:"templates"`
}

type Service interface {
	CreateCampaign(c context.Context, req *CreateCampaignReq) (*CreateCampaignRes, error)
	GetCampaigns(c context.Context, req *GetCampaignsReq) (*GetCampaignsRes, error)
	CreateCampaignWithProspects(c context.Context, req *CreateCampaignWithProspectsReq) (*CreateCampaignRes, error)
	ToggleTemplateAssociation(c context.Context, req *ToggleTemplateAssociationReq) (*ToggleTemplateAssociationRes, error)
}
