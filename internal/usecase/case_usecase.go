package usecase

import (
	"context"
	"fmt"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
	"gh5-backend/pkg/utils/trxmanager"
)

type CaseUsecase struct {
	RepositoryFactory repository.Factory
}

func NewCaseUsecase(f repository.Factory) *CaseUsecase {
	return &CaseUsecase{f}
}

func (u *CaseUsecase) Find(ctx context.Context) ([]dto.CaseResponse, error) {
	var result []dto.CaseResponse

	Cases, err := u.RepositoryFactory.CaseRepository.Find(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find all cases : %+v", err)
		return result, err
	}

	lawyers, err := u.RepositoryFactory.UserRepository.FindLawyers(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find contributor by id : %+v", err)
		return result, err
	}

	fmt.Println(lawyers)

	lawyerMap := make(map[string]model.UserModel)
	for _, lawyer := range lawyers {
		lawyerMap[lawyer.ID] = lawyer
	}

	for _, Case := range Cases {
		contributorID := *Case.ContributorID
		uploaderID := *Case.UploaderID

		uploader := lawyerMap[uploaderID]
		contributor := lawyerMap[contributorID]
		Case.Uploader = &uploader
		Case.Contributor = &contributor
		result = append(result, dto.CaseResponse{
			Data: Case,
		})
	}

	return result, nil
}

func (u *CaseUsecase) FindByID(ctx context.Context, payload dto.ByIDRequest) (dto.CaseResponse, error) {
	var result dto.CaseResponse

	data, err := u.RepositoryFactory.CaseRepository.FindByID(ctx, payload.ID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find case by id : %+v", err)
		return result, err
	}

	userContributor, err := u.RepositoryFactory.UserRepository.FindByID(ctx, data.ContributorID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find contributor by id : %+v", err)
		return result, err
	}

	userUploader, err := u.RepositoryFactory.UserRepository.FindByID(ctx, data.UploaderID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find uploader by id : %+v", err)
		return result, err
	}

	data.Contributor = userContributor
	data.Uploader = userUploader

	result = dto.CaseResponse{
		Data: *data,
	}

	return result, nil
}

func (u *CaseUsecase) Create(ctx context.Context, payload dto.CreateCaseRequest) (dto.CaseResponse, error) {
	var (
		result   dto.CaseResponse
		data     model.CaseModel
		caseData = model.CaseModel{
			CaseEntity: model.CaseEntity{
				CaseName:        payload.CaseName,
				CaseType:        payload.CaseType,
				CaseNumber:      payload.CaseNumber,
				CaseDescription: payload.CaseDescription,
				CaseDetail:      payload.CaseDetail,
				Status:          payload.Status,
				UploaderID:      payload.UploaderID,
			},
			Context: ctx,
		}
	)

	data, err := u.RepositoryFactory.CaseRepository.Create(ctx, caseData)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create case : %+v", err)
		return result, err
	}

	result = dto.CaseResponse{
		Data: data,
	}

	return result, nil
}
func (u *CaseUsecase) UpdateByID(ctx context.Context, payload dto.UpdateCaseRequest) (result dto.CaseResponse, err error) {
	var data model.CaseModel

	if err := trxmanager.New(u.RepositoryFactory.Db).WithTrx(ctx, func(ctx context.Context) error {
		existingData, err := u.RepositoryFactory.CaseRepository.FindByID(ctx, payload.ID)
		if err != nil {
			return err
		}

		if payload.CaseName != "" {
			existingData.CaseName = payload.CaseName
		}

		if payload.CaseType != "" {
			existingData.CaseType = payload.CaseType
		}

		if payload.CaseNumber != "" {
			existingData.CaseNumber = payload.CaseNumber
		}

		if payload.CaseNumber != "" {
			existingData.CaseNumber = payload.CaseNumber
		}

		if payload.CaseDetail != "" {
			existingData.CaseDetail = payload.CaseDetail
		}

		if payload.Status != "" {
			existingData.Status = payload.Status
		}

		if payload.IsActive != nil {
			existingData.IsActive = *payload.IsActive
		}

		if payload.UploaderID != nil {
			existingData.UploaderID = payload.UploaderID
		}

		if payload.ContributorID != nil {
			existingData.ContributorID = payload.ContributorID
		}

		data, err = u.RepositoryFactory.CaseRepository.UpdatesByID(ctx, payload.ID, existingData)
		if err != nil {
			return err
		}

		data = *existingData

		return nil
	}); err != nil {
		return result, err
	}

	result = dto.CaseResponse{
		Data: data,
	}

	return result, nil
}
