package dto

import (
	model "gh5-backend/internal/model/entity"
	res "gh5-backend/pkg/utils/response"
)

// request
type (
	AuthLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	AuthRegisterRequest struct {
		model.UserEntity
	}
)

// response
type (
	AuthLoginResponse struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Token    string `json:"token"`
	}
	AuthLoginResponseDoc struct {
		Body struct {
			Meta res.Meta          `json:"meta"`
			Data AuthLoginResponse `json:"data"`
		} `json:"body"`
	}

	AuthRegisterResponse struct {
		model.UserModel
	}
	AuthRegisterResponseDoc struct {
		Body struct {
			Meta res.Meta             `json:"meta"`
			Data AuthRegisterResponse `json:"data"`
		} `json:"body"`
	}
)
