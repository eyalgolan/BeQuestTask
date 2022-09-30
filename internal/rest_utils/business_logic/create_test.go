package business_logic

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/test_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"net/http"
	"testing"
)

type inputAndExpectedResult struct {
	inputAnswer        rest_utils.Answer
	expectedStatusCode int
	expectedErr        error
}

func TestCreate(t *testing.T) {
	basicInputAnswer := rest_utils.Answer{
		Event: "create",
		Data: rest_utils.AnswerData{
			Key:   "name",
			Value: "test",
		},
	}
	var tests = []struct {
		inputsAndResults []inputAndExpectedResult
	}{
		{
			[]inputAndExpectedResult{{
				inputAnswer:        basicInputAnswer,
				expectedStatusCode: http.StatusCreated,
				expectedErr:        nil,
			}},
		},
		{
			[]inputAndExpectedResult{{
				inputAnswer:        basicInputAnswer,
				expectedStatusCode: http.StatusCreated,
				expectedErr:        nil,
			}, {
				inputAnswer:        basicInputAnswer,
				expectedStatusCode: http.StatusBadRequest,
				expectedErr:        ErrDuplicateAnswer,
			},
			},
		},
	}
	for _, tt := range tests {
		dbClient, sqlConnection, err := test_utils.Setup()
		if err != nil {
			err = test_utils.TearDown(dbClient, sqlConnection)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatal(err)
		}

		testName := fmt.Sprintf("\n%+v\n", tt.inputsAndResults)
		t.Run(testName, func(t *testing.T) {
			for _, inputAndResult := range tt.inputsAndResults {
				resultStatusCode, resultErr := Create(inputAndResult.inputAnswer, *dbClient)
				if !errors.Is(resultErr, inputAndResult.expectedErr) {
					t.Errorf("got %+v, want %+v", resultErr, inputAndResult.expectedErr)
				}
				if resultStatusCode != inputAndResult.expectedStatusCode {
					t.Errorf("got %d, want %d", resultStatusCode, inputAndResult.expectedStatusCode)
				}
			}
		})
		err = test_utils.TearDown(dbClient, sqlConnection)
		if err != nil {
			t.Fatal(err)
		}
	}
}
