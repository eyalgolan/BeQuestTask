package business_logic

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/test_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	db, err := test_utils.ConnectToDB() //todo make this part of setup of each test
	if err != nil {
		t.Fatal(err)
	}
	sqlInstance, err := db.DB.DB()
	if err != nil {
		t.Fatal(err)
	}
	defer sqlInstance.Close()
	var tests = []struct {
		inputAnswer        rest_utils.Answer
		expectedStatusCode int
		expectedErr        error
	}{
		{rest_utils.Answer{
			Event: "create",
			Data: rest_utils.AnswerData{
				Key:   "name",
				Value: "test",
			},
		}, http.StatusCreated,
			nil,
		},
	}
	//todo make this part of teardown of each test
	defer t.Cleanup(func() {
		db.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Answer{}, &models.Event{})
	})
	for _, tt := range tests {
		testName := fmt.Sprintf("answer: \n%+v\n expected status code: \n%d\n expected error: \n%+v\n",
			tt.inputAnswer,
			tt.expectedStatusCode,
			tt.expectedErr)
		t.Run(testName, func(t *testing.T) {
			resultStatusCode, resultErr := Create(tt.inputAnswer, *db)
			if !errors.Is(resultErr, tt.expectedErr) {
				t.Errorf("got %+v, want %+v", resultErr, tt.expectedErr)
			}
			if resultStatusCode != http.StatusCreated {
				t.Errorf("got %d, want %d", resultStatusCode, tt.expectedStatusCode)
			}
		})
	}
}
