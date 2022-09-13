package rest_utils

type Answer struct {
	Event string     `json:"event"`
	Data  AnswerData `json:"data"`
}

type AnswerData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
