package test_utils

import (
	"database/sql"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Setup() (*postgress_utils.DBClient, *sql.DB, error) {
	dbClient, err := ConnectToDB()
	if err != nil {
		return nil, nil, err
	}
	sqlConnection, err := dbClient.DB.DB()
	if err != nil {
		return nil, nil, err
	}
	return dbClient, sqlConnection, nil
}

func TearDown(dbClient *postgress_utils.DBClient, sqlInstance *sql.DB) error {
	dbClient.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Answer{}, &models.Event{})
	return sqlInstance.Close()
}

func ConnectToDB() (*postgress_utils.DBClient, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "open db")
	}
	err = db.AutoMigrate(&models.Answer{}, &models.Event{})
	if err != nil {
		return nil, errors.Wrap(err, "migrate db")
	}
	return &postgress_utils.DBClient{DB: db}, nil
}
