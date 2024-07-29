package service

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd/internal/application/dto"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service/command"
	"github.com/Chengxufeng1994/go-ddd/internal/application/service/query"
	"github.com/Chengxufeng1994/go-ddd/internal/application/usecase"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/entity"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/repository"
	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
)

type UserService struct {
	userRepository    repository.UserRepository
	accountRepository repository.AccountRepository
}

func NewUserService(customerRepository repository.UserRepository, accountRepository repository.AccountRepository) usecase.UserUseCase {
	return &UserService{
		userRepository:    customerRepository,
		accountRepository: accountRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	email, _ := valueobject.NewEmail(req.Email)
	entity := &entity.User{
		Email:          email,
		HashedPassword: req.Password,
		UserInfo:       valueobject.NewUserInfo(req.Age, req.FirstName, req.LastName),
	}

	ruser, err := s.userRepository.CreateUser(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &command.CreateUserCommandResult{
		ID:        ruser.ID,
		Active:    ruser.Active,
		Email:     ruser.Email.String(),
		Age:       ruser.UserInfo.Age(),
		FirstName: ruser.UserInfo.FirstName(),
		LastName:  ruser.UserInfo.LastName(),
		CreatedAt: ruser.CreatedAt,
		UpdatedAt: ruser.UpdatedAt,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req query.GetUserQuery) (*query.GetUserQueryResult, error) {
	ruser, err := s.userRepository.GetUser(ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	return &query.GetUserQueryResult{
		ID:        ruser.ID,
		Active:    ruser.Active,
		Email:     ruser.Email.String(),
		Age:       ruser.UserInfo.Age(),
		FirstName: ruser.UserInfo.FirstName(),
		LastName:  ruser.UserInfo.LastName(),
		CreatedAt: ruser.CreatedAt,
		UpdatedAt: ruser.UpdatedAt,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, queryParams *dto.UserQueryParams) (*[]dto.User, *repository.PaginationResult, error) {
	criteria := repository.PaginationCriteria{
		Page:  queryParams.CurrentPage,
		Limit: queryParams.PageSize,
	}
	res, pRes, err := s.userRepository.ListUsers(ctx, criteria)
	if err != nil {
		return nil, nil, err
	}
	dtos := make([]dto.User, 0, len(*res))
	for _, item := range *res {
		dtos = append(dtos, dto.User{
			ID:        item.ID,
			Active:    item.Active,
			Email:     item.Email.String(),
			Age:       item.UserInfo.Age(),
			FirstName: item.UserInfo.FirstName(),
			LastName:  item.UserInfo.LastName(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return &dtos, pRes, nil
}

func (s *UserService) SearchUsers(ctx context.Context, queryParams *dto.UserQueryParams) (*[]dto.User, *repository.PaginationResult, error) {
	criteria := repository.UserSearchCriteria{
		PaginationCriteria: repository.PaginationCriteria{
			Page:  queryParams.CurrentPage,
			Limit: queryParams.PageSize,
		},
		OrderByCriteria: repository.OrderByCriteria{
			SortBy:  queryParams.SortBy,
			OrderBy: queryParams.OrderBy,
		},
		Email: queryParams.Email,
	}

	res, pRes, err := s.userRepository.SearchUsers(ctx, criteria)
	if err != nil {
		return nil, nil, err
	}

	dtos := make([]dto.User, 0, len(*res))
	for _, item := range *res {
		dtos = append(dtos, dto.User{
			ID:        item.ID,
			Active:    item.Active,
			Email:     item.Email.String(),
			Age:       item.UserInfo.Age(),
			FirstName: item.UserInfo.FirstName(),
			LastName:  item.UserInfo.LastName(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return &dtos, pRes, nil
}

func (s *UserService) AddAccountWithUser(ctx context.Context, req *dto.AccountCreationRequest) (*dto.AccountCreationResponse, error) {
	existed, err := s.userRepository.GetUser(ctx, req.UserID)
	if existed == nil || err != nil {
		return nil, err
	}

	money, _ := valueobject.NewMoney(req.Amount, req.Currency)
	entity := &entity.Account{
		UserID: req.UserID,
		Money:  money,
	}

	raccount, err := s.accountRepository.CreateAccount(ctx, entity)
	if err != nil {
		return nil, err
	}

	return &dto.AccountCreationResponse{
		ID:        raccount.ID,
		UserID:    raccount.UserID,
		Amount:    raccount.Money.Amount(),
		Currency:  raccount.Money.Currency(),
		CreatedAt: raccount.CreatedAt,
		UpdatedAt: raccount.UpdatedAt,
	}, nil
}
