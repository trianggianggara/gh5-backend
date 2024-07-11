package dto

import (
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"
)

// request
type (
	CreateUserRequest struct {
		Name     string  `json:"name" validate:"required"`
		Username string  `json:"username" validate:"required"`
		Email    *string `json:"email,omitempty" validate:"omitempty"`
		Password string  `json:"password"`
		RoleID   string  `json:"role_id" validate:"required"`
		Address  string  `json:"address" validate:"required"`
	}
)

type (
	UpdateUserRequest struct {
		ID       string `param:"id" validate:"required"`
		Name     string `json:"name"`
		Email    string `json:"email" validate:"omitempty,email"`
		Password string `json:"password"`
	}
)

type FindByEmailRequest struct {
	Email string `param:"email" validate:"required"`
}

// response
type (
	UserResponse struct {
		Data model.UserModel
	}
	UserResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data UserResponse `json:"data"`
		} `json:"body"`
	}
)
