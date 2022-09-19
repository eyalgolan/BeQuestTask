package db_utils

import "BeQuestPrep/internal/rest_utils"

type UpdateAPI interface {
	UpdateAnswer(answer rest_utils.AnswerData) error
}

func UpdateAnswer(api UpdateAPI, answer rest_utils.AnswerData) error {
	return api.UpdateAnswer(answer)
}
