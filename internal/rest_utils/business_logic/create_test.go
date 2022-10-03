package business_logic

import (
	"fmt"
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/test_utils"
	"github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"
	"github.com/pkg/errors"
	"net/http"
	"testing"
)

func TestCreateBasicCases(t *testing.T) {
	var tests = []struct {
		inputAnswer        rest_utils.Answer
		expectedStatusCode int
		expectedErr        error
	}{
		{test_utils.BasicCreateAnswer,
			http.StatusCreated,
			nil,
		},
		{
			test_utils.EmptyCreateAnswer,
			http.StatusCreated,
			nil,
		},
		{
			test_utils.InvalidEventAnswer,
			http.StatusBadRequest,
			ErrNotCreate,
		},
	}
	for _, tt := range tests {
		testName := fmt.Sprintf("answer: \n%+v\n expected status code: \n%d\n expected error: \n%+v\n",
			tt.inputAnswer,
			tt.expectedStatusCode,
			tt.expectedErr)
		dbClient, sqlConnection, err := test_utils.Setup()
		if err != nil {
			err = test_utils.TearDown(dbClient, sqlConnection)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatal(err)
		}
		t.Run(testName, func(t *testing.T) {
			resultStatusCode, resultErr := Create(tt.inputAnswer, *dbClient)
			if !errors.Is(resultErr, tt.expectedErr) {
				t.Errorf("got %+v, want %+v", resultErr, tt.expectedErr)
			}
			if resultStatusCode != tt.expectedStatusCode {
				t.Errorf("got %d, want %d", resultStatusCode, tt.expectedStatusCode)
			}
		})
		err = test_utils.TearDown(dbClient, sqlConnection)
		if err != nil {
			t.Fatal(err)
		}
	}
}

type inputAndExpectedResult struct {
	inputAnswer        rest_utils.Answer
	expectedStatusCode int
	expectedErr        error
}

func TestCreateDuplicate(t *testing.T) {
	var tests = []struct {
		inputsAndResults []inputAndExpectedResult
	}{
		{
			[]inputAndExpectedResult{{
				inputAnswer:        test_utils.BasicCreateAnswer,
				expectedStatusCode: http.StatusCreated,
				expectedErr:        nil,
			}, {
				inputAnswer:        test_utils.BasicCreateAnswer,
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
