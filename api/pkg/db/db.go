package db

import (
	"fmt"

	"github.com/Sigma-Ratings/sigma-code-challenges/api/config"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConnection(config *config.DatabaseConfig) (db *gorm.DB, err error) {
	connectionParams := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Name,
		config.Password,
	)
	logrus.WithField("params", config).Info("connecting to database")
	db, err = gorm.Open(postgres.Open(connectionParams), &gorm.Config{
		Logger: logger.Default.LogMode(config.LogLevel),
	})
	if err != nil {
		return nil, err
	}
	err = initMigrate(db)
	return db, nil
}

func initMigrate(db *gorm.DB) error {
	//db.Exec("any raw sql script you need to do")
	db.Exec("CREATE INDEX index_double_meta_alias ON sanction_entities USING GIST (dmetaphone(alias) text_pattern_ops);")
	db.Exec("CREATE EXTENSION fuzzystrmatch;")

	if err := db.AutoMigrate(
		&model.SanctionEntity{},
	); err != nil {
		return err
	}

	return nil
}
