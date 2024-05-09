package http

import (
	"fmt"
	"net/http"

	"github.com/Chengxufeng1994/go-ddd/config"
	docs "github.com/Chengxufeng1994/go-ddd/docs"
	"github.com/Chengxufeng1994/go-ddd/internal/adapter/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(controller *controller.Controller) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiv1 := router.Group("/api/v1")
	{
		apiv1.GET("/hello", controller.HelloController.SayHello)
		apiv1.GET("/accounts/:account_id", controller.AccountController.Get)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}

func NewHttpServer(cfg *config.Transport, engine *gin.Engine) *http.Server {
	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	return srv
}
