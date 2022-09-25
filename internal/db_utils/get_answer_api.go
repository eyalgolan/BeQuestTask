package db_utils

import "KeyValuePermStore/internal/db_utils/postgress_utils/models"

type GetAnswerAPI interface {
	GetAnswer(key string) (*models.Answer, error)
}

func GetAnswer(api GetAnswerAPI, key string) (*models.Answer, error) {
	return api.GetAnswer(key)
}
