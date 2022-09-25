package db_utils

import "BeQuestPrep/internal/rest_utils"

type CreateHistoryAPI interface {
	CreateHistory(answer rest_utils.AnswerData) error
}

func CreateHistory(api CreateHistoryAPI, answer rest_utils.AnswerData) error {
	return api.CreateHistory(answer)
}
