package entity

import (
	"context"

	abstraction "gh5-backend/internal/model/base"
	constant "gh5-backend/pkg/constants"
	"gh5-backend/pkg/ctxval"
	"gh5-backend/pkg/utils/date"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleEntity struct {
	Name     string `json:"name" validate:"required" gorm:"size:200; not null"`
	RoleCode string `json:"email" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type RoleModel struct {
	abstraction.Entity
	RoleEntity
	Context context.Context `json:"-" gorm:"-"`
}

func (RoleModel) TableName() string {
	return "roles"
}

func (m *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	m.IsActive = true
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *RoleModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()

	authCtx := ctxval.GetAuthValue(m.Context)
	if authCtx != nil {
		m.ModifiedBy = &authCtx.Name
	}
	return
}
