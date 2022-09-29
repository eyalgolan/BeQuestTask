package business_logic

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	ErrNotCreate       = errors.New("POST request can only have a create event")
	ErrDuplicateAnswer = errors.New("duplicate answer")
)

func Create(answer rest_utils.Answer, db postgress_utils.DBClient) (int, error) {
	if answer.Event != "create" {
		return http.StatusBadRequest, ErrNotCreate
	}
	latestAnswer, err := db.GetLatestAnswer(answer.Data.Key)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusInternalServerError, err
	}
	if latestAnswer != nil {
		if latestAnswer.Value == answer.Data.Value {
			return http.StatusBadRequest, ErrDuplicateAnswer
		}
	}
	err = db.CreateAnswer(answer.Data)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	err = db.CreateEvent(answer.Data)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
