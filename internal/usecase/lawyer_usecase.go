package usecase

import (
	"context"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
	"gh5-backend/pkg/utils/trxmanager"
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

func (u *LawyerUsecase) UpdateByID(ctx context.Context, payload dto.UpdateLawyerRequest) (result dto.LawyerResponse, err error) {
	var data model.LawyerModel

	if err := trxmanager.New(u.RepositoryFactory.Db).WithTrx(ctx, func(ctx context.Context) error {
		existingData, err := u.RepositoryFactory.LawyerRepository.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		if payload.Specialization != nil {
			existingData.Specialization = *payload.Specialization
		}

		if payload.Position != nil {
			existingData.Position = *payload.Position
		}

		data, err = u.RepositoryFactory.LawyerRepository.UpdatesByID(ctx, payload.ID, existingData)
		if err != nil {
			return err
		}

		data = *existingData

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.LawyerResponse{
		Data: data,
	}

	return result, nil
}
