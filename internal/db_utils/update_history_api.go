package db_utils

import "BeQuestPrep/internal/rest_utils"

type UpdateHistoryAPI interface {
	UpdateHistory(answer rest_utils.AnswerData) error
}

func UpdateHistory(api UpdateHistoryAPI, answer rest_utils.AnswerData) error {
	return api.UpdateHistory(answer)
}
