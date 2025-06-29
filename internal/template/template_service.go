package template

import (
	"context"
	"fmt"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/models"
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

func (s *service) CreateTemplate(c context.Context, req *CreateTemplateReq) (*CreateTemplateRes, error) {
	if req.Name == "" {
		var templateCount int64
		s.DB.Where("user_id = ?", req.UserID).Count(&templateCount)
		req.Name = fmt.Sprintf("Template %03d", templateCount+1)
	}

	var template = models.Template{
		Name:    req.Name,
		Subject: req.Subject,
		Body:    req.Content,
		Sender:  req.Sender,
		UserID:  req.UserID,
	}
	s.DB.Create(&template)

	return &CreateTemplateRes{template.ID}, nil
}
