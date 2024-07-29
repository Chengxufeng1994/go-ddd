package controller

import (
	"net/http"
	"strconv"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/dto/common"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service/query"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService usecase.UserUseCase
}

func NewUserController(userService usecase.UserUseCase) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (ctrl *UserController) Query(c *gin.Context) {
	var req dto.UserQueryParams
	if err := c.ShouldBindQuery(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	res, pRes, err := ctrl.userService.SearchUsers(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, common.PageResponse{
		Code: http.StatusOK,
		Data: &common.PageResult{
			PaginationResult: pRes,
			Rows:             res,
		},
		Msg: "success",
	})
}

func (ctrl *UserController) Get(c *gin.Context) {
	val := c.Param("id")
	id, err := strconv.Atoi(val)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	res, err := ctrl.userService.GetUser(c.Request.Context(), query.GetUserQuery{UserID: uint(id)})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Code: http.StatusOK,
		Data: res,
		Msg:  "success",
	})
}
