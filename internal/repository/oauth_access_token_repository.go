package oauth

import "gorm.io/gorm"

type OauthAccessTokenRepository interface {
	Create(oauthAccessToken entity.OauthAccessToken) (*entity.OauthAccessToken, error)
	Delete(id int) error
	
}

type OauthAccessTokenImpl struct {
	db *gorm.DB
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository{
	return &OauthAccessTokenImpl{db}

}