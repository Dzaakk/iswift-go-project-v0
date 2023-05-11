//go:build wireinject
// +build wireinject

package oauth

import (
	oauthHandler "iswift-go-project/internal/oauth/delivery/http"
	oauthRepository "iswift-go-project/internal/oauth/repository"
	oauthUseCase "iswift-go-project/internal/oauth/usecase"
	userRepository "iswift-go-project/internal/user/repository"
	userUseCase "iswift-go-project/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthHandler.NewOauthHandler,
		oauthRepository.NewOauthClientRepository,
		oauthRepository.NewOauthAccessTokenRepository,
		oauthRepository.NewOauthRefreshTokenRepository,
		oauthUseCase.NewOauthUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &oauthHandler.OauthHandler{}
}
