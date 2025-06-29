package email

import "context"

type SyncEMailsReq struct {
	UserID uint
}

type SyncEMailsResp struct{}

type Service interface {
	SyncEMails(c context.Context, req *SyncEMailsReq) (*SyncEMailsResp, error)
}
