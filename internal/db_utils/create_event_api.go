package db_utils

import "BeQuestPrep/internal/rest_utils"

type CreateEventAPI interface {
	CreateEvent(answer rest_utils.AnswerData) error
}

func CreateEvent(api CreateEventAPI, answer rest_utils.AnswerData) error {
	return api.CreateEvent(answer)
}