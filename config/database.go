package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func NewDatabase(config Configuration) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", config.DatabaseConnection)
	defer db.Close()

	if err != nil {
		return nil, err
	}

	logrus.Info("successful database connection")
	return db, nil
}
