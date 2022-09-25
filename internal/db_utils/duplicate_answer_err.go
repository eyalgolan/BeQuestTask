package db_utils

import (
	"KeyValuePermStore/internal/rest_utils"
	"fmt"
)

type DuplicateAnswerErr struct {
	Answer rest_utils.AnswerData
}

func (e *DuplicateAnswerErr) Error() string {
	return fmt.Sprintf("duplicate answer: %+v", e.Answer)
}
