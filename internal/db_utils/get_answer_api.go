package db_utils

import "github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"

type GetAnswerAPI interface {
	GetAnswer(key string) (*models.Answer, error)
	GetLatestAnswer(key string) (*models.Answer, error)
}

func GetAnswer(api GetAnswerAPI, key string) (*models.Answer, error) {
	return api.GetAnswer(key)
}

func GetLatestAnswer(api GetAnswerAPI, key string) (*models.Answer, error) {
	return api.GetLatestAnswer(key)
}
