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
	dsn := "host=10.1.5.7 user=postgres password=P@ssw0rd dbname=postgres port=31820 sslmode=disable TimeZone=UTC"
	db, err := persistence.New(dsn)
	if err != nil {
		fmt.Printf("error connecting database: %s\n", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&po.Customer{}, &po.Account{}, &po.Transfer{})
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
	customerService := service.NewCustomerService(customerRepository, accountRepository)
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

	// ctx := context.Background()
	// customerCreationResp1, _ := customerService.CreateCustomer(ctx, &dto.CustomerCreationRequest{
	// 	Email:     "goddd@example.com",
	// 	Password:  "password",
	// 	Age:       30,
	// 	FirstName: "first",
	// 	LastName:  "last",
	// })
	// fmt.Println(customerCreationResp1)

	// accountCreationResp1, _ := customerService.AddAccountWithCustomer(ctx, &dto.AccountCreationRequest{
	// 	CustomerID: customerCreationResp1.ID,
	// 	Amount:     1000,
	// 	Currency:   "USD",
	// })
	// fmt.Println(accountCreationResp1)
	// customerCreationResp2, _ := customerService.CreateCustomer(ctx, &dto.CustomerCreationRequest{
	// 	Email:     "goddd_2@example.com",
	// 	Password:  "password",
	// 	Age:       30,
	// 	FirstName: "first",
	// 	LastName:  "last",
	// })
	// fmt.Println(customerCreationResp1)

	// accountCreationResp2, _ := customerService.AddAccountWithCustomer(ctx, &dto.AccountCreationRequest{
	// 	CustomerID: customerCreationResp2.ID,
	// 	Amount:     1000,
	// 	Currency:   "USD",
	// })
	// fmt.Println(accountCreationResp2)
	// customers, err := customerRepository.ListCustomers(ctx, intrepo.PaginationCriteria{Page: 2, Limit: 1})
	// fmt.Println(customers)

	// res2, _ := accountService.GetAccount(ctx, 1)
	// res3, _ := accountService.GetAccount(ctx, 2)
	// transactionService.TransferWithTrx(ctx, &dto.TransferRequest{
	// 	FromAccountId: res2.ID,
	// 	ToAccountId:   res3.ID,
	// 	Amount:        100,
	// })
}
