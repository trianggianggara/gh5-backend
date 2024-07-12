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

type VoteEntity struct {
	IsActive bool    `json:"is_active"`
	UserID   *string `json:"user_id"`
	CaseID   *string `json:"case_id"`

	Users *UserModel `json:"user" gorm:"foreignkey:UserID"`
	Cases *CaseModel `json:"case" gorm:"foreignkey:CaseID"`
}

type VoteModel struct {
	abstraction.Entity
	VoteEntity
	Context context.Context `json:"-" gorm:"-"`
}

func (VoteModel) TableName() string {
	return "votes"
}

func (m *VoteModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	m.IsActive = true
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *VoteModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()

	authCtx := ctxval.GetAuthValue(m.Context)
	if authCtx != nil {
		m.ModifiedBy = &authCtx.Name
	}
	return
}

type VoteCount struct {
	CaseID    string `json:"case_id"`
	VoteCount int    `json:"vote_count"`

	Case CaseModel `json:"case" gorm:"foreignkey:CaseID"`
}
