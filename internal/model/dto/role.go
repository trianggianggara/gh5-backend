package dto

import (
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"
)

// request
type (
	CreateRoleRequest struct {
		model.RoleEntity
	}
)

type (
	UpdateRoleRequest struct {
		ID       string `param:"id" validate:"required"`
		Name     string `json:"name" validate:"required"`
		RoleCode string `json:"role_code" validate:"required"`
	}
)

// response
type (
	RoleResponse struct {
		Data model.RoleModel
	}
	RoleResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data RoleResponse `json:"data"`
		} `json:"body"`
	}
)
