package business_logic

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

func Get(key string, db postgress_utils.DBClient) (int, *rest_utils.AnswerData, error) {
	if key == "" {
		return http.StatusBadRequest, nil, errors.New("GET request must include answer key")
	}
	answer, err := db.GetLatestAnswer(key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, nil, errors.New("Not found")
		}
		return http.StatusInternalServerError, nil, errors.New("unable to perform request")
	}
	return http.StatusOK, &rest_utils.AnswerData{Key: answer.Key, Value: answer.Value}, nil
}
