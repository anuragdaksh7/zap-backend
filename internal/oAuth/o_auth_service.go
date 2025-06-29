package oAuth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"time"

	"github.com/anuragdaksh7/zapmail-backend/config"
	"github.com/anuragdaksh7/zapmail-backend/logger"
	"github.com/anuragdaksh7/zapmail-backend/models"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
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
	url := config.GoogleOAuthConfig.AuthCodeURL(
		"random-state-token",
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("prompt", "consent"),
	)
	return &GetRedirectURLRes{
		RedirectURL: url,
	}, nil
}

func (s *service) HandleGoogleCallback(c context.Context, req *HandleGoogleCallbackReq) (*HandleGoogleCallbackRes, error) {
	_config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), req.Code)
	if err != nil {
		logger.Logger.Error("Failed to exchange token", zap.Error(err))
		return nil, errors.New("unable to exchange token")
	}

	client := config.GoogleOAuthConfig.Client(context.Background(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		logger.Logger.Error("Failed to get userinfo", zap.Error(err))
		return nil, errors.New("unable to get userinfo")
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var googleUser GoogleUser
	if err := json.Unmarshal(body, &googleUser); err != nil {
		logger.Logger.Error("Failed to unmarshal userinfo", zap.Error(err))
	}

	var user models.User
	s.DB.Where("email = ?", googleUser.Email).First(&user)
	if user.ID == 0 {
		logger.Logger.Info("user not found", zap.String("email", googleUser.Email))
		user = models.User{
			Name:       googleUser.Name,
			Email:      googleUser.Email,
			Password:   "",
			IsVerified: true,
			OTP:        "",
			ProfilePic: googleUser.Picture,
		}
		s.DB.Create(&user)
	}

	var userToken models.UserToken
	s.DB.Where("user_id = ?", user.ID).First(&userToken)
	if userToken.ID == 0 {
		logger.Logger.Info("no existing token found, create one")
		userToken = models.UserToken{
			UserID:             user.ID,
			GoogleAccessToken:  token.AccessToken,
			GoogleRefreshToken: token.RefreshToken,
		}
		s.DB.Create(&userToken)
	} else {
		logger.Logger.Info("existing token found")
		userToken.GoogleAccessToken = token.AccessToken
		userToken.GoogleRefreshToken = token.RefreshToken
		s.DB.Save(&userToken)
	}
	logger.Logger.Info("Token updated")

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(_config.JwtSecret))
	if err != nil {
		return nil, err
	}

	return &HandleGoogleCallbackRes{
		Token:          tokenString,
		ID:             1,
		Name:           user.Name,
		Email:          user.Email,
		ProfilePicture: user.ProfilePic,
	}, nil
}
