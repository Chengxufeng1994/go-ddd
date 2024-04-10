package application

import (
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
)

type ApplicationConfiguration struct {
	AccountService     usecase.AccountUseCase
	CustomerService    usecase.CustomerUseCase
	TransactionService usecase.TransactionUseCase
}

type Application struct {
	AccountService     usecase.AccountUseCase
	CustomerService    usecase.CustomerUseCase
	TransactionService usecase.TransactionUseCase
}

func NewApplication(appCfg *ApplicationConfiguration) *Application {
	return &Application{
		AccountService:     appCfg.AccountService,
		CustomerService:    appCfg.CustomerService,
		TransactionService: appCfg.TransactionService,
	}
}
