package db_utils

import "BeQuestPrep/internal/rest_utils"

type UpdateEventAPI interface {
	UpdateEvent(answer rest_utils.AnswerData) error
}

func UpdateEvent(api UpdateEventAPI, answer rest_utils.AnswerData) error {
	return api.UpdateEvent(answer)
}
