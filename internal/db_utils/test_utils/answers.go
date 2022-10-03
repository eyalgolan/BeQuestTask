package test_utils

import "github.com/eyalgolan/key-value-persistent-store/internal/rest_utils"

var (
	BasicCreateAnswer = rest_utils.Answer{
		Event: "create",
		Data: rest_utils.AnswerData{
			Key:   "name",
			Value: "test",
		},
	}
	EmptyCreateAnswer = rest_utils.Answer{
		Event: "create",
		Data: rest_utils.AnswerData{
			Key:   "",
			Value: "",
		},
	}
	BasicUpdateAnswer = rest_utils.Answer{
		Event: "update",
		Data: rest_utils.AnswerData{
			Key:   "name",
			Value: "test",
		},
	}
	EmptyUpdateAnswer = rest_utils.Answer{
		Event: "update",
		Data: rest_utils.AnswerData{
			Key:   "",
			Value: "",
		},
	}
	InvalidEventAnswer = rest_utils.Answer{
		Event: "something",
		Data: rest_utils.AnswerData{
			Key:   "",
			Value: "",
		},
	}
)
