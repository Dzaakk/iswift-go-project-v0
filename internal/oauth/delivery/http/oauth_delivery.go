package oauth

import (
	dto "iswift-go-project/internal/oauth/dto"
	usecase "iswift-go-project/internal/oauth/usecase"
	"iswift-go-project/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	usecase usecase.OauthUseCase
}

func NewOauthHandler(usecase usecase.OauthUseCase) *OauthHandler {
	return &OauthHandler{usecase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")

	oauthRouter.POST("/oauth", handler.Login)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	// Memanggil usecase dari login
	data, err := handler.usecase.Login(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))

}
