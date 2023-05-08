package oauth

import ("gorm.io/gorm"
	oauth "iswift-go-project/internal/oauth/entity"
)
type OauthClientRepository interface {
}

type OauthClientImpl struct {
	db *gorm.DB
}

// // FindByClientIdAndClientSecret implements OauthClientRepository
// if err := oc.db.Where("client_id = ?", client_id).Where("client_secret = ?",)

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &OauthClientImpl{db}
}

