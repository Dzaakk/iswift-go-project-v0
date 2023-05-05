package register

import (
	dto "iswift-go-project/internal/register/dto"
	userUseCase "iswift-go-project/internal/user/usecase"	
	userDto "iswift-go-project/internal/user/dto"
	mail "iswift-go-project/pkg/mail/sendgrid"
)

type RegisterUseCase interface {
	Register(registerDto dto.CreateRegisterRequestBody) error
}

type RegisterUseCaseImpl struct {
	//

}

func NewRegisterUseCase(userUseCase userUseCase.UserUseCase) RegisterUseCase {
	userUseCase useuserUseCase.UserUseCase
//
	return &RegisterUseCaseImpl{userUseCase}
}

func (ru *RegisterUseCaseImpl) Register(userDto userDto.UserRequestBody) error {
	// Check ke dalam module user
	user, err := ru.userUseCase.Create(userDto)

	
}
