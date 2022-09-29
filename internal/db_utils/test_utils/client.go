package test_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

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
