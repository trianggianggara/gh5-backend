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

func (u *VoteUsecase) VouteCount(ctx context.Context) ([]dto.VoteCountResponse, error) {
	var result []dto.VoteCountResponse

	votes, err := u.RepositoryFactory.VoteRepository.VoteCount(ctx)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find all votes : %+v", err)
		return result, err
	}

	for _, vote := range votes {
		result = append(result, dto.VoteCountResponse{
			Data: vote,
		})
	}

	return result, nil
}

func (u *VoteUsecase) VoteCountByCaseID(ctx context.Context, payload dto.ByIDRequest) (dto.VoteCountResponse, error) {
	var result dto.VoteCountResponse

	data, err := u.RepositoryFactory.VoteRepository.VoteCountByCaseID(ctx, payload.ID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find vote by id : %+v", err)
		return result, err
	}

	result = dto.VoteCountResponse{
		Data: *data,
	}

	return result, nil
}

func (u *VoteUsecase) Vote(ctx context.Context, payload dto.UpvoteRequest) (dto.VoteResponse, error) {
	var (
		result   dto.VoteResponse
		data     model.VoteModel
		voteData = model.VoteModel{
			VoteEntity: model.VoteEntity{
				CaseID: payload.CaseID,
				UserID: payload.UserID,
			},
			Context: ctx,
		}
	)

	vote, err := u.RepositoryFactory.VoteRepository.FindVoteByCaseAndUserID(ctx, payload.CaseID, payload.UserID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed find vote by case and user id : %+v", err)
		return result, err
	}

	if vote.ID != "" {
		if vote.IsActive {
			data, err = u.RepositoryFactory.VoteRepository.Downvote(ctx, payload.CaseID, payload.UserID)
			if err != nil {
				u.RepositoryFactory.Log.Warnf("Failed to revote : %+v", err)
				return result, err
			}
		} else {
			data, err = u.RepositoryFactory.VoteRepository.Revote(ctx, payload.CaseID, payload.UserID)
			if err != nil {
				u.RepositoryFactory.Log.Warnf("Failed create Vote : %+v", err)
				return result, err
			}
		}
	} else {
		data, err = u.RepositoryFactory.VoteRepository.Create(ctx, voteData)
		if err != nil {
			u.RepositoryFactory.Log.Warnf("Failed create Vote : %+v", err)
			return result, err
		}
	}

	result = dto.VoteResponse{
		Data: data,
	}

	return result, nil
}

func (u *VoteUsecase) Downvote(ctx context.Context, caseID string, userID string) (dto.VoteResponse, error) {
	var result dto.VoteResponse

	data, err := u.RepositoryFactory.VoteRepository.Downvote(ctx, caseID, userID)
	if err != nil {
		u.RepositoryFactory.Log.Warnf("Failed create Vote : %+v", err)
		return result, err
	}

	result = dto.VoteResponse{
		Data: data,
	}

	return result, nil
}
