package oAuth

import (
	"context"
	"net/http"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"golang.org/x/oauth2"
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

func (s *service) GetRedirectURL(c context.Context) (*GetRedirectURLRes, error) {
	url := config.GoogleOAuthConfig.AuthCodeURL("random-state-token", oauth2.AccessTypeOffline)
	return &GetRedirectURLRes{
		RedirectURL: url,
	}, nil
}

func (s *service) GetClient(c context.Context) (*http.Client, error) {
	return &http.Client{Timeout: s.timeout}, nil
}
