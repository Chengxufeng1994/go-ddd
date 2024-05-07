package main

import (
	"errors"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Chengxufeng1994/go-ddd/internal/adapter/controller"
	"github.com/Chengxufeng1994/go-ddd/internal/adapter/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/application"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http"
)

func main() {
	dsn := "host=localhost user=root password=P@ssw0rd dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	db, err := persistence.New(dsn)
	if err != nil {
		fmt.Printf("error connecting database: %s\n", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&po.User{}, &po.Account{}, &po.Transfer{})
	if err != nil {
		fmt.Printf("error migrating database: %s\n", err)
		os.Exit(1)
	}

	errc := make(chan error, 1)

	customerRepository := repository.NewGormCustomerRepository(db)
	accountRepository := repository.NewGormAccountRepository(db)
	transferRepository := repository.NewGormTransferRepository(db)
	unitOfWorkRepository := repository.New(db)

	accountService := service.NewAccountService(accountRepository, customerRepository)
	customerService := service.NewUserService(customerRepository, accountRepository)
	transactionService := service.NewTransactionService(transferRepository, accountRepository, unitOfWorkRepository)

	appCfg := &application.ApplicationConfiguration{
		AccountService:     accountService,
		CustomerService:    customerService,
		TransactionService: transactionService,
	}
	app := application.NewApplication(appCfg)

	ctrl := controller.NewController(app)

	router := http.NewRouter(ctrl)

	httpSrv := http.NewHttpServer(router)

	go func() {
		err := httpSrv.ListenAndServe()
		if errors.Is(err, nethttp.ErrServerClosed) {
			fmt.Printf("server one closed\n")
			errc <- err
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
			os.Exit(1)
		}
	}()

	go func() {
		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		// kill (no param) default send syscanll.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("received signal %v", <-quit)
	}()

	er := <-errc
	fmt.Printf("exit: %v\n", er)
}
