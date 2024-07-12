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

type CaseEntity struct {
	CaseName        string  `json:"case_name" gorm:"not null"`
	CaseType        string  `json:"case_type"`
	CaseNumber      string  `json:"case_number" gorm:"size:50"`
	CaseDescription string  `json:"case_description" gorm:"not null"`
	CaseDetail      string  `json:"case_detail"`
	Status          string  `json:"status" validate:"required" gorm:"not null"`
	IsActive        bool    `json:"is_active"`
	ContributorID   *string `json:"contributor_id"`
	UploaderID      *string `json:"uploader_id"`

	Contributor *UserModel `json:"contributor" gorm:"foreignkey:ContributorID"`
	Uploader    *UserModel `json:"uploader" gorm:"foreignkey:UploaderID"`
}

type CaseModel struct {
	abstraction.Entity
	CaseEntity
	Context context.Context `json:"-" gorm:"-"`
}

func (CaseModel) TableName() string {
	return "cases"
}

func (m *CaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	m.IsActive = true
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *CaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()

	authCtx := ctxval.GetAuthValue(m.Context)
	if authCtx != nil {
		m.ModifiedBy = &authCtx.Name
	}
	return
}
