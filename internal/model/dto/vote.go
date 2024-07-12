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

type (
	VoteCountResponse struct {
		Data model.VoteCount
	}
	VoteCountResponseDoc struct {
		Body struct {
			Meta res.Meta          `json:"meta"`
			Data VoteCountResponse `json:"data"`
		} `json:"body"`
	}
)
