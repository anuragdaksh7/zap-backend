package oAuth

import "context"

type GetRedirectURLRes struct {
	RedirectURL string `json:"redirect_url"`
}

type HandleGoogleCallbackReq struct {
	Code string `json:"code"`
}

type HandleGoogleCallbackRes struct {
	Token          string `json:"token"`
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profilePicture"`
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
}

type Service interface {
	GetRedirectURL(c context.Context) (*GetRedirectURLRes, error)
	HandleGoogleCallback(c context.Context, req *HandleGoogleCallbackReq) (*HandleGoogleCallbackRes, error)
}
