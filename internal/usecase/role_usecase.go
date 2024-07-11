package usecase

import (
	"context"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
)

type RoleUsecase struct {
	RepositoryFactory repository.Factory
}

func NewRoleUsecase(f repository.Factory) *RoleUsecase {
	return &RoleUsecase{f}
}

func (u *RoleUsecase) Find(ctx context.Context) ([]dto.RoleResponse, error) {
	var result []dto.RoleResponse

	roles, err := u.RepositoryFactory.RoleRepository.Find(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find all roles : %+v", err)
		return result, err
	}

	for _, role := range roles {
		result = append(result, dto.RoleResponse{
			Data: role,
		})
	}

	return result, nil
}

func (u *RoleUsecase) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.RoleResponse, error) {
	var result dto.RoleResponse

	data, err := u.RepositoryFactory.RoleRepository.FindByID(ctx, payload.ID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find role by id : %+v", err)
		return result, err
	}

	result = dto.RoleResponse{
		Data: *data,
	}

	return result, nil
}

func (u *RoleUsecase) Create(ctx context.Context, payload dto.CreateRoleRequest) (dto.RoleResponse, error) {
	var (
		result dto.RoleResponse
		data   model.RoleModel
		Role   = model.RoleModel{
			RoleEntity: model.RoleEntity{
				Name:     payload.Name,
				RoleCode: payload.RoleCode,
			},
			Context: ctx,
		}
	)

	data, err := u.RepositoryFactory.RoleRepository.Create(ctx, Role)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create role : %+v", err)
		return result, err
	}

	result = dto.RoleResponse{
		Data: data,
	}

	return result, nil
}
