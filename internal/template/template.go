package template

import "context"

type CreateTemplateReq struct {
	Sender  string `json:"sender"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Name    string `json:"name"`
	UserID  uint   `json:"user_id"`
}

type CreateTemplateRes struct {
	ID uint `json:"id"`
}

type Service interface {
	CreateTemplate(c context.Context, req *CreateTemplateReq) (*CreateTemplateRes, error)
}
