package dto

import (
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"
)

// request
type (
	UpvoteRequest struct {
		model.VoteEntity
	}
)

type (
	DownvoteRequest struct {
		UserID string `queryparam:"user_id"`
		CaseID string `queryparam:"case_id"`
	}
)

// response
type (
	VoteResponse struct {
		Data model.VoteModel
	}
	VoteResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data VoteResponse `json:"data"`
		} `json:"body"`
	}
)
