package controller

import (
	"net/http"
	"strconv"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/dto/common"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuService usecase.MenuUseCase
}

func NewMenuController(menuService usecase.MenuUseCase) *MenuController {
	return &MenuController{
		menuService: menuService,
	}
}

func (ctrl *MenuController) Create(c *gin.Context) {
	var req dto.MenuCreationRequest
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	menu, err := ctrl.menuService.CreateMenu(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		Code: common.SUCCESS,
		Data: menu,
		Msg:  "success",
	})
}

func (ctrl *MenuController) Get(c *gin.Context) {
	val := c.Param("id")
	id, err := strconv.Atoi(val)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, common.Response{
			Code: http.StatusBadRequest,
			Msg:  "bad request",
		})
		return
	}

	res, err := ctrl.menuService.GetMenu(c.Request.Context(), uint(id))
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

func (ctrl *MenuController) Update(c *gin.Context) {}

func (ctrl *MenuController) Delete(c *gin.Context) {}
