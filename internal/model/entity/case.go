package entity

import (
	"context"
	"time"

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
	Document        string  `json:"document"`
	Status          string  `json:"status" gorm:"not null"`
	IsActive        bool    `json:"is_active"`
	ClientID        *string `json:"client_id"`
	ContributorID   *string `json:"contributor_id"`
	UploaderID      *string `json:"uploader_id"`

	Client      *UserModel `json:"client" gorm:"foreignkey:ClientID"`
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
	m.Status = "Pending"
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

type CaseDetails struct {
	CaseID          string    `json:"case_id"`
	CreatedAt       time.Time `json:"created_at"`
	CaseNumber      string    `json:"case_number"`
	CaseDescription string    `json:"case_description"`
	CaseDetail      string    `json:"case_detail"`
	Status          string    `json:"status"`
	CaseName        string    `json:"case_name"`
	CaseType        string    `json:"case_type"`
	ClientID        string    `json:"client_id"`
	ContributorID   string    `json:"contributor_id"`
	UploaderID      string    `json:"uploader_id"`
	UserID          string    `json:"user_id"`

	Client      *UserModel `json:"client" gorm:"foreignkey:ClientID"`
	Contributor *UserModel `json:"contributor" gorm:"foreignkey:ContributorID"`
	Uploader    *UserModel `json:"uploader" gorm:"foreignkey:UploaderID"`
}

func (CaseDetails) TableName() string {
	return "case_details"
}
