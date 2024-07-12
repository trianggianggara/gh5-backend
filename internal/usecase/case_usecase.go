package usecase

import (
	"context"
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

	lawyerUsers, err := u.RepositoryFactory.UserRepository.FindLawyers(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find lawyer users id : %+v", err)
		return result, err
	}

	lawyers, err := u.RepositoryFactory.LawyerRepository.Find(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find layers: %+v", err)
		return result, err
	}

	lawyerMap := make(map[string]model.LawyerModel)
	for _, lawyer := range lawyers {
		lawyerMap[lawyer.ID] = lawyer
	}

	lawyerUserMap := make(map[string]model.UserModel)
	for _, lawyerUser := range lawyerUsers {
		lawyerUserMap[lawyerUser.ID] = lawyerUser
	}

	for _, Case := range Cases {
		var (
			contributorID string
			uploaderID    string
		)

		if Case.ContributorID != nil {
			contributorID = *Case.ContributorID

			contributor := lawyerUserMap[contributorID]
			Case.Contributor = &contributor

			lawyer := lawyerMap[*contributor.LawyerID]
			Case.Contributor.Lawyer = &lawyer
		}

		if Case.UploaderID != nil {
			uploaderID = *Case.UploaderID

			uploader := lawyerUserMap[uploaderID]
			Case.Uploader = &uploader

			lawyer := lawyerMap[*uploader.LawyerID]
			Case.Uploader.Lawyer = &lawyer

		}

		client, err := u.RepositoryFactory.UserRepository.FindByID(ctx, Case.ClientID)
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed find lawyer by id : %+v", err)
			return result, err
		}

		Case.Client = client

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

	if data.ContributorID != nil {
		userContributor, err := u.RepositoryFactory.UserRepository.FindByID(ctx, data.ContributorID)
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed find contributor by id : %+v", err)
			return result, err
		}
		data.Contributor = userContributor

		lawyer, err := u.RepositoryFactory.LawyerRepository.FindByID(ctx, data.Contributor.LawyerID)
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed find lawyer by id : %+v", err)
			return result, err
		}

		data.Contributor.Lawyer = lawyer
	}

	if data.UploaderID != nil {
		userUploader, err := u.RepositoryFactory.UserRepository.FindByID(ctx, data.UploaderID)
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed find uploader by id : %+v", err)
			return result, err
		}
		data.Uploader = userUploader

		lawyer, err := u.RepositoryFactory.LawyerRepository.FindByID(ctx, data.Uploader.LawyerID)
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed find lawyer by id : %+v", err)
			return result, err
		}

		data.Uploader.Lawyer = lawyer
	}

	client, err := u.RepositoryFactory.UserRepository.FindByID(ctx, data.ClientID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find lawyer by id : %+v", err)
		return result, err
	}

	data.Client = client

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
				UploaderID:      payload.UploaderID,
				ClientID:        payload.ClientID,
				ContributorID:   payload.ContributorID,
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
