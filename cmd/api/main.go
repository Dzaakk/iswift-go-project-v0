package main

import (
	registerHandler "iswift-go-project/internal/register/handler"
	registerUseCase "iswift-go-project/internal/register/usecase"
	userRepository "iswift-go-project/internal/user/repository"
	userUseCase "iswift-go-project/internal/user/usecase"
	mysql "iswift-go-project/pkg/db/mysql"
	sendgrid "iswift-go-project/pkg/mail/sendgrid"
	"github.com/gin-gonic/gin"
)

func main() {
	db := mysql.DB()

	r := gin.Default()
	mail := sendgrid.NewMail()
	userRepository := userRepository.NewUserRepositoryImpl(db)
	userUseCasse := userUseCase.NewUserUseCase(userRepository)
	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCasse, mail)
	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)
	r.Run()
	
}

