package usecase

import (
	"context"
	"gh5-backend/internal/factory/repository"
	"gh5-backend/internal/model/dto"
	model "gh5-backend/internal/model/entity"
)

type VoteUsecase struct {
	RepositoryFactory repository.Factory
}

func NewVoteUsecase(f repository.Factory) *VoteUsecase {
	return &VoteUsecase{f}
}

func (u *VoteUsecase) Upvote(ctx context.Context, payload dto.UpvoteRequest) (dto.VoteResponse, error) {
	var (
		result dto.VoteResponse
		data   model.VoteModel
		Vote   = model.VoteModel{
			VoteEntity: model.VoteEntity{
				CaseID: payload.CaseID,
				UserID: payload.UserID,
			},
			Context: ctx,
		}
	)

	data, err := u.RepositoryFactory.VoteRepository.Create(ctx, Vote)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create Vote : %+v", err)
		return result, err
	}

	result = dto.VoteResponse{
		Data: data,
	}

	return result, nil
}

func (u *VoteUsecase) Downvote(ctx context.Context, payload dto.DownvoteRequest) (dto.VoteResponse, error) {
	var result dto.VoteResponse

	data, err := u.RepositoryFactory.VoteRepository.Downvote(ctx, payload.CaseID, payload.UserID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create Vote : %+v", err)
		return result, err
	}

	result = dto.VoteResponse{
		Data: data,
	}

	return result, nil
}
