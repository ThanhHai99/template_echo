package services

import (
	"Template_Echo/pkg/configs"
	"Template_Echo/pkg/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type DatabaseConfig struct {
	IP           string `env:"DB_HOST" envDefault:"localhost"`
	Port         int    `env:"DB_PORT" envDefault:"5432"`
	DatabaseName string `env:"DB_NAME" envDefault:"echo"`
	User         string `env:"DB_USER" envDefault:"root"`
	Password     string `env:"DB_PASS" envDefault:"root"`
	GormPreload  bool   `env:"DB_PRELOAD" envDefault:"true"`
}

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		var e0 error
		config := DatabaseConfig{}
		if config.GormPreload {
			db = db.Set("gorm:auto_preload", true)
			utils.Log().Info().Msg("Set auto preload mode")
		}

		connectionString := fmt.Sprintf("postgresql://%s@%s:%d/%s?sslmode=disable", config.User, config.IP, config.Port, config.DatabaseName)
		db, e0 = gorm.Open("postgres", connectionString)
		if e0 != nil {
			utils.Log().Error().Err(e0).Msg("failed to connect database")
			utils.Log().Debug().Msg("reconnecting to database...")
			time.Sleep(5 * time.Second)
			return DB()
		}

		utils.Log().Info().Msg("Connected to database")
		if configs.DbIsDebug() {
			db.LogMode(true)
			utils.Log().Info().Msg("Set database on debuging mode")
		}
		AddUUIDGenerateExtension(db)
		AddUnaccentExtension(db)
		if config.GormPreload {
			db = db.Set("gorm:auto_preload", true)
			utils.Log().Info().Msg("Set auto preload mode")
		}
	}

	if err := db.DB().Ping(); err != nil {
		utils.Log().Error().Err(err).Msg("ping failed to database")
		utils.Log().Debug().Msg("reconnecting to database...")
		_ = db.Close()
		db = nil
		return DB()
	}

	return db
}

func AddUUIDGenerateExtension(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		utils.Log().Fatal().Msg("Can't install extension uuid-ossp")
	}
	utils.Log().Info().Msg("Add uuid-ossp extension")
}

func AddUnaccentExtension(db *gorm.DB) {
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "unaccent"`).Error; err != nil {
		utils.Log().Fatal().Msg("Can't install extension unaccent")
	}
	utils.Log().Info().Msg("Add unaccent extension")
}
