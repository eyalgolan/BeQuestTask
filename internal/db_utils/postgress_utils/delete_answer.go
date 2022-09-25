package postgress_utils

import (
	"KeyValuePermStore/internal/rest_utils"
	"github.com/pkg/errors"
)

func (c *Client) DeleteAnswer(answer rest_utils.AnswerData) error {
	latestAnswer, err := c.getLatestAnswer(answer.Key)
	latestAnswer.Deleted = true
	err = c.DB.Save(&latestAnswer).Error
	if err != nil {
		return errors.Wrap(err, "delete answer")
	}
	return nil
}
