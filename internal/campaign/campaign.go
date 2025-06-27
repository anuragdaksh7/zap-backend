package campaign

import (
	"context"
	"time"
)

type CreateCampaignReq struct {
	UserID           uint   `json:"user_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	RunEveryNMinutes uint   `json:"run_every_n_minutes"`
	TotalProspects   uint   `json:"total_prospects"`
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

type Service interface {
	CreateCampaign(c context.Context, req *CreateCampaignReq) (*CreateCampaignRes, error)
}
