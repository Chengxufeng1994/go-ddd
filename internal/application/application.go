package application

import (
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
)

type ApplicationConfiguration struct {
	AuthService        usecase.AuthUseCase
	UserService        usecase.UserUseCase
	AccountService     usecase.AccountUseCase
	TransactionService usecase.TransactionUseCase
}

type Application struct {
	AuthService        usecase.AuthUseCase
	UserService        usecase.UserUseCase
	AccountService     usecase.AccountUseCase
	TransactionService usecase.TransactionUseCase
}

func NewApplication(appCfg *ApplicationConfiguration) *Application {
	return &Application{
		AuthService:        appCfg.AuthService,
		UserService:        appCfg.UserService,
		AccountService:     appCfg.AccountService,
		TransactionService: appCfg.TransactionService,
	}
}
