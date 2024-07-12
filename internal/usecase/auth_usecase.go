package usecase

import (
	"context"

	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"

	"gh5-backend/pkg/utils/trxmanager"

	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	RepositoryFactory repository.Factory
}

func NewAuthUsecase(f repository.Factory) *AuthUsecase {
	return &AuthUsecase{f}
}

func (s *AuthUsecase) Login(ctx context.Context, payload dto.AuthLoginRequest) (dto.AuthLoginResponse, error) {
	var result dto.AuthLoginResponse

	data, err := s.RepositoryFactory.UserRepository.FindByEmail(ctx, payload.Email)
	if data == nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.EmailOrPasswordIncorrect, err)
	}

	token, err := data.GenerateJWT()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = dto.AuthLoginResponse{
		Name:     data.Name,
		Username: data.Username,
		Token:    token,
	}

	return result, nil
}

func (s *AuthUsecase) Register(ctx context.Context, payload dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {
	var result dto.AuthRegisterResponse
	var data model.UserModel
	var err error

	data.UserEntity = payload.UserEntity

	if err = trxmanager.New(s.RepositoryFactory.Db).WithTrx(ctx, func(ctx context.Context) error {
		data, err = s.RepositoryFactory.UserRepository.Create(ctx, data)
		if err != nil {
			return err
		}

		role, err := s.RepositoryFactory.RoleRepository.FindByID(ctx, data.RoleID)
		if err != nil {
			return err
		}

		if role.RoleCode == "LYR" {
			lawyer := model.LawyerModel{
				LawyerEntity: model.LawyerEntity{},
				Context:      ctx,
			}
			_, err = s.RepositoryFactory.LawyerRepository.Create(ctx, lawyer)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return result, err
	}

	result = dto.AuthRegisterResponse{
		UserModel: data,
	}

	return result, nil
}
