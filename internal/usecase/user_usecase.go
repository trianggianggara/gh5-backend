package usecase

import (
	"context"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
	"gh5-backend/pkg/utils/str"
)

type UserUsecase struct {
	RepositoryFactory repository.Factory
}

func NewUserUsecase(f repository.Factory) *UserUsecase {
	return &UserUsecase{f}
}

func (u *UserUsecase) Find(ctx context.Context) ([]dto.UserResponse, error) {
	var result []dto.UserResponse

	users, err := u.RepositoryFactory.UserRepository.Find(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find all users: %+v", err)
		return result, err
	}

	roles, err := u.RepositoryFactory.RoleRepository.Find(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find all roles: %+v", err)
		return result, err
	}

	roleMap := make(map[string]model.RoleModel)
	for _, role := range roles {
		roleMap[role.ID] = role
	}

	for _, user := range users {
		roleData := roleMap[user.RoleID]
		user.Role = &roleData
		result = append(result, dto.UserResponse{
			Data: user,
		})
	}

	return result, nil
}

func (u *UserUsecase) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.UserResponse, error) {
	var result dto.UserResponse

	data, err := u.RepositoryFactory.UserRepository.FindByID(ctx, payload.ID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find user by id : %+v", err)
		return result, err
	}

	role, err := u.RepositoryFactory.RoleRepository.FindByID(ctx, data.RoleID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find all roles: %+v", err)
		return result, err
	}

	data.Role = role

	result = dto.UserResponse{
		Data: *data,
	}

	return result, nil
}

func (u *UserUsecase) FindByEmail(ctx context.Context, payload dto.FindByEmailRequest) (dto.UserResponse, error) {
	var result dto.UserResponse

	data, err := u.RepositoryFactory.UserRepository.FindByEmail(ctx, payload.Email)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find user by email : %+v", err)
		return result, err
	}

	result = dto.UserResponse{
		Data: *data,
	}

	return result, nil
}

func (u *UserUsecase) Create(ctx context.Context, payload dto.CreateUserRequest) (dto.UserResponse, error) {
	var (
		result dto.UserResponse
		email  string
	)

	if payload.Email != nil {
		email = *payload.Email
	} else {
		email = str.GenerateRandString(10) + "@gmail.com"
	}

	var (
		data model.UserModel
		user = model.UserModel{
			UserEntity: model.UserEntity{
				Email:    email,
				Username: payload.Username,
				Password: payload.Password,
				RoleID:   payload.RoleID,
			},
			Context: ctx,
		}
	)

	data, err := u.RepositoryFactory.UserRepository.Create(ctx, user)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create user : %+v", err)
		return result, err
	}

	result = dto.UserResponse{
		Data: data,
	}

	return result, nil
}
