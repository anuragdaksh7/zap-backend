package _google

import (
	"context"

	"github.com/anuragdaksh7/zapmail-backend/logger"
	"go.uber.org/zap"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Gmail interface {
	List(ctx context.Context) (error, error)
}

type _gmail struct {
	client option.ClientOption
}

func NewGmail(client option.ClientOption) Gmail {
	return &_gmail{client}
}

func (g *_gmail) List(ctx context.Context) (error, error) {
	user := "me"
	srv, err := gmail.NewService(ctx, g.client)
	if err != nil {
		logger.Logger.Error("Error creating gmail service", zap.Error(err))
		return nil, err
	}
	logger.Logger.Info("Gmail service created")

	messageList, err := srv.Users.Messages.List(user).MaxResults(50).Do()
	if err != nil {
		logger.Logger.Error("Error retrieving messages", zap.Error(err))
		return nil, err
	}
	logger.Logger.Info("Messages retrieved")
	for _, m := range messageList.Messages {
		logger.Logger.Info("Message", zap.String("Id", m.Id))
	}

	return nil, nil
}
