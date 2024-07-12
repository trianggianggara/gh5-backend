package usecase

import (
	"context"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
	"gh5-backend/pkg/document"
	"gh5-backend/pkg/utils/trxmanager"

	"github.com/google/uuid"
)

type CaseUsecase struct {
	RepositoryFactory repository.Factory
}

func NewCaseUsecase(f repository.Factory) *CaseUsecase {
	return &CaseUsecase{f}
}

func (u *CaseUsecase) Find(ctx context.Context) ([]dto.CaseResponse, error) {
	var result []dto.CaseResponse

	cases, err := u.RepositoryFactory.CaseRepository.FindAll(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed to find all cases: %+v", err)
		return result, err
	}

	for _, c := range cases {
		result = append(result, dto.CaseResponse{
			Data: c,
		})
	}

	return result, nil
}

func (u *CaseUsecase) FindByUserID(ctx context.Context, userID string) ([]dto.CaseDetailsResponse, error) {
	var result []dto.CaseDetailsResponse

	cases, err := u.RepositoryFactory.CaseRepository.FindCaseByUserID(ctx, userID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed to find all cases: %+v", err)
		return result, err
	}

	for _, c := range cases {
		result = append(result, dto.CaseDetailsResponse{
			Data: c,
		})
	}

	return result, nil
}

func (u *CaseUsecase) FindByLawyerID(ctx context.Context, lawyerID string, status string) ([]dto.LawyerCaseResponse, error) {
	var result []dto.LawyerCaseResponse

	cases, err := u.RepositoryFactory.CaseRepository.FindCaseByLawyerID(ctx, lawyerID, status)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed to find all cases: %+v", err)
		return result, err
	}

	for _, c := range cases {
		result = append(result, dto.LawyerCaseResponse{
			Data: c,
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

	if payload.Document != nil {
		docUrl, err := document.UploadAndSavePath(ctx, payload.Document, "case_docs", uuid.NewString())
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed to upload document: %+v", err)
			return result, err
		}
		caseData.Document = docUrl
	}

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

		if payload.Document != nil {
			docUrl, err := document.UploadAndSavePath(ctx, payload.Document, "case_docs", uuid.NewString())
			if err != nil {
				u.RepositoryFactory.Log.Warnf("Failed to upload document: %+v", err)
				return err
			}
			existingData.Document = docUrl
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
