package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloController struct{}

func NewHelloController() *HelloController {
	return &HelloController{}
}

//	@BasePath	/api/v1

// PingExample godoc
//
//	@Summary	ping example
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	helloworld
//	@Router			/hello [get]
func (ctrl *HelloController) SayHello(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
	return
}
