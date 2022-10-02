package business_logic

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	ErrNotUpdate = errors.New("PUT request can only have an update event")
	ErrNoKey     = errors.New("key doesn't exist")
)

func Update(answer rest_utils.Answer, db postgress_utils.DBClient) (int, *rest_utils.Answer, error) {
	if answer.Event != "update" {
		return http.StatusBadRequest, nil, ErrNotUpdate
	}
	err := db.UpdateAnswer(answer.Data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, nil, ErrNoKey
		}
		return http.StatusInternalServerError, nil, err
	}
	err = db.UpdateEvent(answer.Data)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, &answer, nil
}
