package controller

import "github.com/Chengxufeng1994/go-ddd/internal/application"

type Controller struct {
	HelloController   *HelloController
	AccountController *AccountController
}

func NewController(app *application.Application) *Controller {
	return &Controller{
		HelloController:   NewHelloController(),
		AccountController: NewAccountController(app.AccountService),
	}
}
