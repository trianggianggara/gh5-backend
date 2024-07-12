package dto

import (
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"
)

// request
type (
	CreateLawyerRequest struct {
		model.LawyerEntity
	}
)

type (
	UpdateLawyerRequest struct {
		ID             string  `param:"id" validate:"required"`
		Position       *string `param:"position"`
		Specialization *string `json:"specialization"`
		IsActive       bool    `json:"is_active"`
	}
)

// response
type (
	LawyerResponse struct {
		Data model.LawyerModel
	}
	LawyerResponseDoc struct {
		Body struct {
			Meta res.Meta       `json:"meta"`
			Data LawyerResponse `json:"data"`
		} `json:"body"`
	}
)
