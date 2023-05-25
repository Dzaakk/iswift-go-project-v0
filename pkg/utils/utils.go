package utils

import (
	"math/rand"

	oauthDto "iswift-go-project/internal/oauth/dto"

	"github.com/gin-gonic/gin"
)

func RandString(number int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, number)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponse {
	user, _ := ctx.Get("user")

	return user.(*oauthDto.MapClaimsResponse)
}
