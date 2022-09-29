package db_utils

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
)

type DuplicateAnswerErr struct {
	Answer rest_utils.AnswerData
}

func (e *DuplicateAnswerErr) Error() string {
	return fmt.Sprintf("duplicate answer: %+v", e.Answer)
}
