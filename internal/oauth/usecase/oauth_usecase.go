package oauth

import (
	"database/sql"
	"errors"
	dto "iswift-go-project/internal/oauth/dto"
	entity "iswift-go-project/internal/oauth/entity"
	repository "iswift-go-project/internal/oauth/repository"
	userUseCase "iswift-go-project/internal/user/usecase"
	utils "iswift-go-project/pkg/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type OauthUseCase interface {
	Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error)
	Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error)
}

type OauthUseCaseImpl struct {
	oauthClientRepository       repository.OauthClientRepository
	oauthAccessTokenRepository  repository.OauthAccessTokenRepository
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository
	userUseCase                 userUseCase.UserUseCase
}

// Login implements OauthUseCase
func (usecase *OauthUseCaseImpl) Login(loginRequestBody dto.LoginRequestBody) (*dto.LoginResponse, error) {
	// check oauth client id dan oauth client secret

	oauthClient, err := usecase.oauthClientRepository.FindByClientIdAndClientSecret(loginRequestBody.ClientId, loginRequestBody.ClientSecret)

	if err != nil {
		return nil, err
	}

	var user dto.UserResponse

	// Login menggunakan user
	dataUser, err := usecase.userUseCase.FindByEmail(loginRequestBody.Email)

	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	user.ID = dataUser.ID
	user.Name = dataUser.Name
	user.Email = dataUser.Email
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	// Compare password apakah sama atau tidak
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestBody.Password))
	if err != nil {
		return nil, errors.New("username or password is invalid")
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &dto.ClaimsResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	// Insert data ke table oauth access token
	dataOauthAccessToken := entity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		Token:         tokenString,
		Scope:         "*",
		ExpiredAt: sql.NullTime{
			Time: expirationTime,
		},
	}

	oauthAccessToken, err := usecase.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	// Insert data ke table oauth refresh token
	dataOauthRefreshToken := entity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		Token:              utils.RandString(128),
		ExpiredAt: sql.NullTime{
			Time: time.Now().Add(24 * 366 * time.Hour),
		},
	}

	oauthRefreshToken, err := usecase.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  tokenString,
		RefreshToken: oauthRefreshToken.Token,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil
}

// Refresh implements OauthUseCase
func (*OauthUseCaseImpl) Refresh(refreshTokenRequestBody dto.RefreshTokenRequestBody) (*dto.LoginResponse, error) {
	panic("unimplemented")
}

func NewOauthUseCase(
	oauthClientRepository repository.OauthClientRepository,
	oauthAccessTokenRepository repository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository repository.OauthRefreshTokenRepository,
	userUseCase userUseCase.UserUseCase,
) OauthUseCase {
	return &OauthUseCaseImpl{oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUseCase}
}