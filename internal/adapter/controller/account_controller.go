package controller

import (
	"net/http"
	"strconv"

	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService usecase.AccountUseCase
}

func NewAccountController(accountService usecase.AccountUseCase) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

// ShowAccount godoc
// @Summary			Show an account
// @Schemes
// @Description	get account by ID
// @Tags				accounts
// @Accept			json
// @Produce			json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  dto.AccountResponse
// @Router			/accounts/{id} [get]
func (ctrl *AccountController) Get(c *gin.Context) {
	val := c.Param("account_id")
	id, err := strconv.Atoi(val)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid account ID",
		})
		return
	}

	resp, err := ctrl.accountService.GetAccount(c.Request.Context(), uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "account not found",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
