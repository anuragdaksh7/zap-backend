package models

import (
	"github.com/anuragdaksh7/zapmail-backend/utils"
	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model
	UserID             uint
	GoogleAccessToken  string
	GoogleRefreshToken string

	User User `gorm:"foreignKey:UserID;references:ID"`
}

func (t *UserToken) BeforeSave(tx *gorm.DB) (err error) {
	t.GoogleAccessToken, err = utils.Encrypt(t.GoogleAccessToken)
	if err != nil {
		return err
	}
	t.GoogleRefreshToken, err = utils.Encrypt(t.GoogleRefreshToken)
	return
}

func (t *UserToken) AfterFind(tx *gorm.DB) (err error) {
	t.GoogleAccessToken, err = utils.Decrypt(t.GoogleAccessToken)
	if err != nil {
		return err
	}
	t.GoogleRefreshToken, err = utils.Decrypt(t.GoogleRefreshToken)
	return
}
