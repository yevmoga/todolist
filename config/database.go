package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

func NewDatabase(config Configuration) *gorm.DB {
	db, err := gorm.Open("postgres", config.DatabaseConnection)

	if err != nil {
		panic("error connection to database: " + err.Error())
	}

	db.LogMode(true)
	db.SetLogger(logrus.StandardLogger())

	logrus.Info("successful database connection")
	return db
}
