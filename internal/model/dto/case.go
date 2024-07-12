package dto

import (
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"
)

// request
type (
	CreateCaseRequest struct {
		model.CaseEntity
	}
)

type (
	UpdateCaseRequest struct {
		ID              string  `param:"id" validate:"required"`
		CaseName        string  `json:"case_name"`
		CaseType        string  `json:"case_type"`
		CaseNumber      string  `json:"case_number"`
		CaseDescription string  `json:"case_description"`
		CaseDetail      string  `json:"case_detail"`
		Status          string  `json:"status" validate:"required"`
		IsActive        *bool   `json:"is_active"`
		ContributorID   *string `json:"contributor_id"`
		UploaderID      *string `json:"uploader_id"`
	}
)

// response
type (
	CaseResponse struct {
		Data model.CaseModel
	}
	CaseResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data CaseResponse `json:"data"`
		} `json:"body"`
	}
)
