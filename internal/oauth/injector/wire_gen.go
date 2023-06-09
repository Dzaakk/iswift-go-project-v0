// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package oauth

import (
	"gorm.io/gorm"
	"iswift-go-project/internal/admin/repository"
	admin2 "iswift-go-project/internal/admin/usecase"
	"iswift-go-project/internal/oauth/delivery/http"
	oauth2 "iswift-go-project/internal/oauth/repository"
	oauth3 "iswift-go-project/internal/oauth/usecase"
	"iswift-go-project/internal/user/repository"
	user2 "iswift-go-project/internal/user/usecase"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *oauth.OauthHandler {
	oauthClientRepository := oauth2.NewOauthClientRepository(db)
	oauthAccessTokenRepository := oauth2.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := oauth2.NewOauthRefreshTokenRepository(db)
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	adminRepository := admin.NewAdminRepository(db)
	adminUseCase := admin2.NewAdminUseCase(adminRepository)
	oauthUseCase := oauth3.NewOauthUseCase(oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUseCase, adminUseCase)
	oauthHandler := oauth.NewOauthHandler(oauthUseCase)
	return oauthHandler
}
