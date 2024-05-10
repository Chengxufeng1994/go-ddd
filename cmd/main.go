package main

import (
	"context"
	"errors"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Chengxufeng1994/go-ddd/config"
	"github.com/Chengxufeng1994/go-ddd/internal/application"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	domainrepository "github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http/controller"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("error load configuration: %s\n", err)
		os.Exit(1)
	}

	db, err := persistence.New(&cfg.Infrastructure.Persistence)
	if err != nil {
		fmt.Printf("error connecting database: %s\n", err)
		os.Exit(1)
	}

	err = db.AutoMigrate(&po.User{}, &po.Role{}, &po.UserRole{}, &po.Permission{}, &po.RolePermission{}, &po.Account{}, &po.Transfer{})
	if err != nil {
		fmt.Printf("error migrating database: %s\n", err)
		os.Exit(1)
	}

	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}
	enforcer, _ := casbin.NewEnforcer("config/model.conf", adapter)
	//add policy
	if hasPolicy, _ := enforcer.HasPolicy("system-admin", "accounts", "write"); !hasPolicy {
		enforcer.AddPolicy("system-admin", "accounts", "write")
	}
	if hasPolicy, _ := enforcer.HasPolicy("system-admin", "accounts", "read"); !hasPolicy {
		enforcer.AddPolicy("system-admin", "accounts", "read")
	}
	if hasPolicy, _ := enforcer.HasPolicy("admin", "accounts", "read"); !hasPolicy {
		enforcer.AddPolicy("admin", "accounts", "read")
	}

	errc := make(chan error, 1)

	userRepository := repository.NewGormUserRepository(db)
	roleRepository := repository.NewGormRoleRepository(db)
	permissionRepository := repository.NewGormPermissionRepository(db)
	// rbacRepository := repository.NewRBACRepository(db)
	accountRepository := repository.NewGormAccountRepository(db)
	transferRepository := repository.NewGormTransferRepository(db)
	unitOfWorkRepository := repository.New(db)
	RoleSeeds(roleRepository)
	PermissionSeeds(permissionRepository)
	AssignPermissionsSeeds(permissionRepository)
	UserSeeds(userRepository)

	accountService := service.NewAccountService(accountRepository, userRepository)
	userService := service.NewUserService(userRepository, accountRepository)
	transactionService := service.NewTransactionService(transferRepository, accountRepository, unitOfWorkRepository)

	appCfg := &application.ApplicationConfiguration{
		AccountService:     accountService,
		UserService:        userService,
		TransactionService: transactionService,
	}
	app := application.NewApplication(appCfg)

	ctrl := controller.NewController(app)
	router := http.NewRouter(enforcer, ctrl)
	httpSrv := http.NewHttpServer(&cfg.Transport, router)

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

func RoleSeeds(repo domainrepository.RoleRepository) {
	roles := []*entity.Role{
		{
			Name: "super-admin",
			Slug: "super-admin",
		},
		{
			Name: "admin",
			Slug: "admin",
		},
		{
			Name: "guest",
			Slug: "guest",
		},
	}

	for _, role := range roles {
		repo.CreateRole(context.Background(), role)
	}
}

func PermissionSeeds(repo domainrepository.PermissionRepository) {
	permissions := []*entity.Permission{
		{
			Name: "GET:Hello",
			Slug: "get hello",
		},
		{
			Name: "GET:Account",
			Slug: "get account",
		},
	}

	for _, perm := range permissions {
		repo.CreatePermission(context.Background(), perm)
	}
}

func AssignPermissionsSeeds(repo domainrepository.PermissionRepository) {
	repo.AssignPermissionsToRole(context.Background(), 1, []uint{1, 2})
	repo.AssignPermissionsToRole(context.Background(), 2, []uint{1})
}

func UserSeeds(repo domainrepository.UserRepository) {

	users := []*entity.User{
		{
			Active:         true,
			Email:          valueobject.MustNewEmail("super_admin@example.com"),
			HashedPassword: "P@ssw0rd",
			UserInfo:       valueobject.NewCustomerInfo(30, "super", "admin"),
			RoleID:         1,
			Roles: []entity.Role{
				{
					ID: 1,
				},
			},
		},
		{
			Active:         true,
			Email:          valueobject.MustNewEmail("guest@example.com"),
			HashedPassword: "P@ssw0rd",
			UserInfo:       valueobject.NewCustomerInfo(30, "guest", "guest"),
			RoleID:         2,
			Roles: []entity.Role{
				{
					ID: 2,
				},
			},
		},
	}

	for _, u := range users {
		repo.CreateUser(context.Background(), u)
	}
}
