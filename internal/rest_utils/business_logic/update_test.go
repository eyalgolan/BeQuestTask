package business_logic

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/test_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"net/http"
	"testing"
)

func TestUpdate(t *testing.T) {
	var tests = []struct {
		updateAnswer       rest_utils.Answer
		createBefore       bool
		expectedStatusCode int
		expectedAnswer     *rest_utils.Answer
		expectedErr        error
	}{
		{test_utils.BasicUpdateAnswer,
			true,
			http.StatusOK,
			&test_utils.BasicUpdateAnswer,
			nil,
		},
		{test_utils.BasicUpdateAnswer,
			false,
			http.StatusBadRequest,
			nil,
			ErrNoKey,
		},
		{
			test_utils.EmptyUpdateAnswer,
			true,
			http.StatusOK,
			&test_utils.EmptyCreateAnswer,
			nil,
		},
		{
			test_utils.InvalidEventAnswer,
			true,
			http.StatusBadRequest,
			nil,
			ErrNotUpdate,
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("answer: \n%+v\n create before: \n%t\nexpected status code: \n%d\n expected answer data: \n%+v\n expected error: \n%+v\n",
			tt.updateAnswer,
			tt.createBefore,
			tt.expectedStatusCode,
			tt.expectedAnswer,
			tt.expectedErr)
		dbClient, sqlConnection, err := test_utils.Setup()
		if err != nil {
			err = test_utils.TearDown(dbClient, sqlConnection)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatal(err)
		}
		if tt.createBefore {
			err = dbClient.CreateAnswer(tt.updateAnswer.Data)
			if err != nil {
				t.Fatal(err)
			}
		}
		t.Run(testName, func(t *testing.T) {
			resultStatusCode, resultAnswer, resultErr := Update(tt.updateAnswer, *dbClient)
			if !errors.Is(resultErr, tt.expectedErr) {
				t.Errorf("got %+v, want %+v", resultErr, tt.expectedErr)
			}
			if resultStatusCode != tt.expectedStatusCode {
				t.Errorf("got %d, want %d", resultStatusCode, tt.expectedStatusCode)
			}
			if tt.expectedAnswer != nil {
				if resultAnswer.Data.Key != tt.expectedAnswer.Data.Key {
					t.Errorf("got %s, want %s", resultAnswer.Data.Key, tt.expectedAnswer.Data.Key)
				}
				if resultAnswer.Data.Value != tt.expectedAnswer.Data.Value {
					t.Errorf("got %s, want %s", resultAnswer.Data.Value, tt.expectedAnswer.Data.Value)
				}
			} else if resultAnswer != nil {
				t.Errorf("got %+v instead of nil", resultAnswer)
			}
		})
		err = test_utils.TearDown(dbClient, sqlConnection)
		if err != nil {
			t.Fatal(err)
		}
	}
}
