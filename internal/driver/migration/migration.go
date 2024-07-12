package migration

import (
	"fmt"

	dbDriver "gh5-backend/internal/driver/db"
	model "gh5-backend/internal/model/entity"
	constant "gh5-backend/pkg/constants"
	"gh5-backend/pkg/utils/env"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	if !env.NewEnv().GetBool(constant.MIGRATION_ENABLED) {
		return
	}

	mgConfigurations := map[string]Migration{
		constant.DB_GH5_BACKEND: &migration{
			DbModels: &[]interface{}{
				&model.UserModel{},
				&model.RoleModel{},
				&model.LawyerModel{},
				&model.CaseModel{},
				&model.VoteModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := dbDriver.GetConnection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}
}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
