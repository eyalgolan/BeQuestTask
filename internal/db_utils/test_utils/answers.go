package test_utils

import "github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"

var (
	BasicAnswer = rest_utils.Answer{
		Event: "create",
		Data: rest_utils.AnswerData{
			Key:   "name",
			Value: "test",
		},
	}
	EmptyAnswer = rest_utils.Answer{
		Event: "create",
		Data: rest_utils.AnswerData{
			Key:   "",
			Value: "",
		},
	}
	NotCreateAnswer = rest_utils.Answer{
		Event: "something",
		Data: rest_utils.AnswerData{
			Key:   "",
			Value: "",
		},
	}
)
