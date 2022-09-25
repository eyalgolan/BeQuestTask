package postgress_utils

import (
	"KeyValuePermStore/internal/db_utils/postgress_utils/models"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Client struct {
	DB *gorm.DB
}

type PostgresConfig struct {
	Host     string `env:"DB_HOST,default=answers_db"`
	User     string `env:"DB_USER,default=gorm"`
	Password string `env:"DB_PASSWORD,default=gorm"`
	DBName   string `env:"DB_NAME,default=answers"`
	Port     int64  `env:"DB_PORT,default=5432"`
	SSLMode  string `env:"SSL_MODE,default=disable"`
}

func ConnectToDB(cfg PostgresConfig) (*Client, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
		cfg.SSLMode)
	time.Sleep(5 * time.Second) // waiting for db to start. TODO change
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf(err.Error())
		return nil, errors.Wrap(err, "open db")
	}
	err = db.AutoMigrate(&models.Answer{}, &models.Event{})
	if err != nil {
		return nil, errors.Wrap(err, "migrate db")
	}
	return &Client{DB: db}, nil
}
