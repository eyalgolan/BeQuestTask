package db_utils

import "KeyValuePermStore/internal/rest_utils"

type DeleteEventAPI interface {
	DeleteEvent(answer rest_utils.AnswerData) error
}

func DeleteEvent(api DeleteEventAPI, answer rest_utils.AnswerData) error {
	return api.DeleteEvent(answer)
}
