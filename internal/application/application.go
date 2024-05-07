package application

import (
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
)

type ApplicationConfiguration struct {
	AccountService     usecase.AccountUseCase
	UserService        usecase.UserUseCase
	TransactionService usecase.TransactionUseCase
}

type Application struct {
	AccountService     usecase.AccountUseCase
	UserService        usecase.UserUseCase
	TransactionService usecase.TransactionUseCase
}

func NewApplication(appCfg *ApplicationConfiguration) *Application {
	return &Application{
		AccountService:     appCfg.AccountService,
		UserService:        appCfg.UserService,
		TransactionService: appCfg.TransactionService,
	}
}
