package business_logic

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	ErrDeleteEmptyKey = errors.New("DELETE request must include answer key")
	ErrKeyMismatch    = errors.New("key param must equal key in body")
	ErrEventNotDelete = errors.New("DELETE request can only have a delete event")
	ErrDeleteNotFound = errors.New("record not found")
)

func Delete(key string, answerToDelete rest_utils.Answer, db postgress_utils.DBClient) (int, *rest_utils.AnswerData, error) {
	if key == "" {
		return http.StatusBadRequest, nil, ErrDeleteEmptyKey
	}
	if answerToDelete.Data.Key != key {
		return http.StatusBadRequest, nil, ErrKeyMismatch
	}
	if answerToDelete.Event != "delete" {
		return http.StatusBadRequest, nil, ErrEventNotDelete
	}
	err := db.DeleteAnswer(answerToDelete.Data)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, nil, ErrDeleteNotFound
		} else {
			return http.StatusInternalServerError, nil, err
		}
	}
	err = db.DeleteEvent(answerToDelete.Data)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, &answerToDelete.Data, nil
}
