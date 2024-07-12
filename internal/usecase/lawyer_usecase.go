package usecase

import (
	"context"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
)

type LawyerUsecase struct {
	RepositoryFactory repository.Factory
}

func NewLawyerUsecase(f repository.Factory) *LawyerUsecase {
	return &LawyerUsecase{f}
}

func (u *LawyerUsecase) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.LawyerResponse, error) {
	var result dto.LawyerResponse

	data, err := u.RepositoryFactory.LawyerRepository.FindByID(ctx, payload.ID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find Lawyer by id : %+v", err)
		return result, err
	}

	result = dto.LawyerResponse{
		Data: *data,
	}

	return result, nil
}

func (u *LawyerUsecase) Create(ctx context.Context, payload dto.CreateLawyerRequest) (dto.LawyerResponse, error) {
	var (
		result dto.LawyerResponse
		data   model.LawyerModel
		Lawyer = model.LawyerModel{
			LawyerEntity: model.LawyerEntity{
				Specialization: payload.Specialization,
			},
			Context: ctx,
		}
	)

	data, err := u.RepositoryFactory.LawyerRepository.Create(ctx, Lawyer)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create Lawyer : %+v", err)
		return result, err
	}

	result = dto.LawyerResponse{
		Data: data,
	}

	return result, nil
}
