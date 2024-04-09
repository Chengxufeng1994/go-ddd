package main

import (
	"context"
	"fmt"

	"github.com/Chengxufeng1994/go-ddd/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/adapter/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence"
	"github.com/Chengxufeng1994/go-ddd/internal/infrastructure/persistence/po"
)

func main() {
	var err error
	dsn := "host=10.1.5.7 user=postgres password=P@ssw0rd dbname=postgres port=31820 sslmode=disable TimeZone=UTC"
	db, err := persistence.New(dsn)
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&po.Customer{})
	if err != nil {
		panic("failed to migrate database")
	}

	req := &dto.CustomerCreationRequest{
		Email:     "goddd@example.com",
		Age:       30,
		FirstName: "first",
		LastName:  "last",
	}
	ctx := context.Background()
	customerRepository := repository.NewGormCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerService.CreateCustomer(ctx, req)
	entities, err := customerRepository.ListCustomers(context.Background())
	fmt.Println(entities)
}
