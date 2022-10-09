package business_logic

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/test_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	var tests = []struct {
		key                string
		expectedStatusCode int
		expectedAnswerData *rest_utils.AnswerData
		expectedErr        error
	}{
		{
			"name",
			http.StatusOK,
			&test_utils.BasicCreateAnswer.Data,
			nil,
		},
		{
			"key that doesn't exist",
			http.StatusBadRequest,
			nil,
			ErrNotFound,
		},
		{
			"",
			http.StatusBadRequest,
			nil,
			ErrGetEmptyKey,
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("key: \n%s\n expected status code: \n%d\n expected answer: \n%+v\n expected error: \n%+v\n",
			tt.key,
			tt.expectedStatusCode,
			tt.expectedAnswerData,
			tt.expectedErr)
		dbClient, sqlConnection, err := test_utils.Setup()
		if err != nil {
			err = test_utils.TearDown(dbClient, sqlConnection)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatal(err)
		}
		if tt.expectedAnswerData != nil {
			err = dbClient.CreateAnswer(*tt.expectedAnswerData)
			if err != nil {
				t.Fatal(err)
			}
		}
		t.Run(testName, func(t *testing.T) {
			resultStatusCode, resultAnswerData, resultErr := Get(tt.key, *dbClient)
			if !errors.Is(resultErr, tt.expectedErr) {
				t.Errorf("got %+v, want %+v", resultErr, tt.expectedErr)
			}
			if resultStatusCode != tt.expectedStatusCode {
				t.Errorf("got %d, want %d", resultStatusCode, tt.expectedStatusCode)
			}
			if tt.expectedAnswerData != nil {
				if resultAnswerData.Key != tt.expectedAnswerData.Key {
					t.Errorf("got %s, want %s", resultAnswerData.Key, tt.expectedAnswerData.Key)
				}
				if resultAnswerData.Value != tt.expectedAnswerData.Value {
					t.Errorf("got %s, want %s", resultAnswerData.Value, tt.expectedAnswerData.Value)
				}
			} else if resultAnswerData != nil {
				t.Errorf("got %+v instead of nil", resultAnswerData)
			}
		})
		err = test_utils.TearDown(dbClient, sqlConnection)
		if err != nil {
			t.Fatal(err)
		}
	}
}
