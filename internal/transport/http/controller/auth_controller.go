package controller

import (
	"net/http"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/dto/common"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService usecase.AuthUseCase
}

func NewAuthController(authService usecase.AuthUseCase) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// SignUp godoc
//
//	@Summary	registers a user
//	@Schemes
//	@Description	registers a user
//	@Tags			authenticate
//	@Accept			json
//	@Produce		json
//	@Param			signUpRequest	body		dto.SignUpRequest	true	"register a user"
//	@Success		200				{object}	common.Response{data=dto.AccountResponse,msg=string}
//	@Failure		404				{object}	common.Response
//	@Router			/signup [post]
func (ctrl *AuthController) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	if err := ctrl.authService.SignUp(c.Request.Context(), &req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Code: common.SUCCESS,
		Msg:  "success",
	})
}

func (ctrl *AuthController) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	if err := ctrl.authService.SignIn(c.Request.Context(), &req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Code: common.SUCCESS,
		Msg:  "success",
	})
}

func (ctrl *AuthController) SignOut(c *gin.Context) {}
