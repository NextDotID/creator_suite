package model

import (
	"github.com/nextdotid/creator_suite/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	l  = logrus.WithFields(logrus.Fields{"module": "model"})
)

// Init initializes DB connection instance and do migration at startup.
func Init() {
	if DB != nil { // initialized
		return
	}
	dsn := config.GetDatabaseDSN()
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		l.Fatalf("Error when opening DB: %s\n", err.Error())
	}

	err = DB.AutoMigrate(
		&Content{},
		&KeyPair{},
	)
	if err != nil {
		panic(err)
	}

	l.Info("database initialized")
}
