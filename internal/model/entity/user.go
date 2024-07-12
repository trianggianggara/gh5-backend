package entity

import (
	"context"
	"os"
	"time"

	abstraction "gh5-backend/internal/model/base"
	constant "gh5-backend/pkg/constants"
	"gh5-backend/pkg/ctxval"
	"gh5-backend/pkg/utils/date"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	Name               string  `json:"name" validate:"required" gorm:"size:200; not null"`
	Username           string  `json:"username" validate:"required" gorm:"size:100;not null"`
	Email              string  `json:"email" validate:"required,email" gorm:"index:idx_user_email;unique;size:150;not null"`
	IdentityCardNumber string  `json:"identity_card_number" validate:"required"`
	Address            string  `json:"address" validate:"required"`
	PasswordHash       string  `json:"-"`
	Password           string  `json:"password" validate:"required" gorm:"-"`
	IsActive           bool    `json:"is_active"`
	RoleID             string  `json:"role_id" validate:"required" gorm:"not null"`
	LawyerID           *string `json:"lawyer_id"`

	Role   *RoleModel   `json:"role" gorm:"foreignkey:RoleID"`
	Lawyer *LawyerModel `json:"lawyer" gorm:"foreignkey:LawyerID"`
}

type UserModel struct {
	abstraction.Entity
	UserEntity
	Context context.Context `json:"-" gorm:"-"`
}

func (UserModel) TableName() string {
	return "users"
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New().String()
	m.IsActive = true
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	m.HashPassword()
	m.Password = ""
	return
}

func (m *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()

	authCtx := ctxval.GetAuthValue(m.Context)
	if authCtx != nil {
		m.ModifiedBy = &authCtx.Name
	}
	return
}

func (m *UserModel) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserModel) GenerateJWT() (string, error) {
	var (
		jwtKey = os.Getenv(constant.JWT_KEY)
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        m.ID,
		"email":     m.Email,
		"name":      m.Name,
		"role":      m.Role.Name,
		"role_code": m.Role.RoleCode,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
