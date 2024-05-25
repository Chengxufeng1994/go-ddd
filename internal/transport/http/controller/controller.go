package controller

import (
	"github.com/Chengxufeng1994/go-ddd/internal/application"
)

type Controller struct {
	HelloController   *HelloController
	AuthController    *AuthController
	UserController    *UserController
	MenuController    *MenuController
	AccountController *AccountController
}

func NewController(app *application.Application) *Controller {
	return &Controller{
		HelloController:   NewHelloController(),
		UserController:    NewUserController(app.UserService),
		MenuController:    NewMenuController(app.MenuService),
		AuthController:    NewAuthController(app.AuthService),
		AccountController: NewAccountController(app.AccountService),
	}
}
