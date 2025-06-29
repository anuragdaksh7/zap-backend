package email

import (
	"context"
	"errors"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/config"
	_google "github.com/anuragdaksh7/zapmail-backend/internal/google"
	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
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

func (s *service) SyncEMails(c context.Context, req *SyncEMailsReq) (*SyncEMailsResp, error) {
	var userToken models.UserToken

	s.DB.Where("user_id = ?", req.UserID).First(&userToken)
	if userToken.ID == 0 {
		logger.Logger.Error("User Token Not Found")
		return nil, errors.New("user token not found")
	}
	if userToken.GoogleAccessToken == "" {
		logger.Logger.Error("User Google Access Token Not Found")
		return nil, errors.New("user google token not found")
	}

	token := &oauth2.Token{
		AccessToken:  userToken.GoogleAccessToken,
		RefreshToken: userToken.GoogleRefreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(-time.Minute),
	}

	tokenSource := oauth2.StaticTokenSource(token)

	gClient := option.WithTokenSource(tokenSource)

	var _gmail = _google.NewGmail(gClient)
	err, err2 := _gmail.List(c)
	if err != nil || err2 != nil {
		return nil, err2
	}

	return nil, nil
}
