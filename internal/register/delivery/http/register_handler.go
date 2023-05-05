package register

import (
	"fmt"
	dto "iswift-go-project/internal/register/dto"
	registerUseCase "iswift-go-project/internal/register/usecase"
	"iswift-go-project/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerUseCase registerUseCase.RegisterUseCase
}

func NewRegisterHandler(registerUseCase registerUseCase.RegisterUseCase) RegisterHandler {
	return RegisterHandler{registerUseCase}
}

func (rh *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/registers", rh.Register)
}

func (rh *RegisterHandler) Register(ctx *gin.Context) {
	// Validate input
	var registerRequestInput dto.CreateRegisterRequestBody

	//validasi dari body yang dikirim menggunakan format json
	if err := ctx.ShouldBindJSON(&registerRequestInput); err != nil {
		fmt.Println("masuk ke dalam error")
		ctx.JSON(http.StatusBadRequest, utils.Response(400, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	err := rh.registerUseCase.Register(registerRequestInput)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(500, "internal server error", err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, utils.Response(201, "created", "succes, please check your email"))
}
