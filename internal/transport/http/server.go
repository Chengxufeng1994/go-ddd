package http

import (
	"fmt"
	"net/http"

	"github.com/Chengxufeng1994/go-ddd/config"
	docs "github.com/Chengxufeng1994/go-ddd/docs"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http/controller"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(enforcer *casbin.Enforcer, controller *controller.Controller) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Ok")
	})
	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.CORSMiddleware())
	{
		apiv1.GET("/hello", controller.HelloController.SayHello)
		apiv1.POST("/signup", controller.AuthController.SignUp)
		apiv1.POST("/signin", controller.AuthController.SignIn)
		apiv1.POST("/signout", controller.AuthController.SignOut)
	}

	userGroup := apiv1.Group("/users")
	{
		userGroup.GET("", controller.UserController.Query)
		userGroup.GET("/:id", controller.UserController.Get)
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
