package http

import (
	"net/http"

	docs "github.com/Chengxufeng1994/go-ddd/docs"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http/controller"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type API struct {
	routes *Routes
	engine *gin.Engine
}

type Routes struct {
	Root  gin.IRouter // ''
	Hello gin.IRouter // "/hello"

	APIv1 gin.IRouter // "api/v1"
	Auth  gin.IRouter // "api/v1/auth"
	Users gin.IRouter // "api/v1/users"

	Menus gin.IRouter // "api/v1/menus"
}

func NewRouter(enforcer *casbin.Enforcer, controller *controller.Controller) *gin.Engine {
	e := gin.Default()
	e.Use(cors.Default())

	api := &API{
		engine: e,
		routes: &Routes{},
	}

	api.routes.Root = api.engine
	api.routes.Hello = api.routes.Root.Group("/hello")
	api.routes.Hello.GET("/", controller.HelloController.SayHello)

	api.routes.APIv1 = api.engine.Group("/api/v1")
	api.routes.APIv1.Use(middleware.CORSMiddleware())
	api.routes.Auth = api.routes.APIv1.Group("/auth")
	api.routes.Users = api.routes.APIv1.Group("/users")
	api.routes.Menus = api.routes.APIv1.Group("/menus")

	api.routes.Auth.POST("/signup", controller.AuthController.SignUp)
	api.routes.Auth.POST("/signin", controller.AuthController.SignIn)
	api.routes.Auth.POST("/signout", controller.AuthController.SignOut)

	api.routes.Users.GET("", controller.UserController.Query)
	api.routes.Users.GET("/:id", controller.UserController.Get)

	api.routes.Menus.POST("", controller.MenuController.Create)
	// 	api.routes.Menus.GET("", controller.MenuController.Query)
	api.routes.Menus.GET("/:id", controller.MenuController.Get)
	api.routes.Menus.PUT("/:id", controller.MenuController.Update)
	api.routes.Menus.DELETE("/:id", controller.MenuController.Delete)

	docs.SwaggerInfo.BasePath = "/api/v1"

	e.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Ok")
	})

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return api.engine
}
