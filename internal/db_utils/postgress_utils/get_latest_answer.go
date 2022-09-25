package postgress_utils

import (
	"KeyValuePermStore/internal/db_utils/postgress_utils/models"
)

func (c *Client) getLatestAnswer(key string) (*models.Answer, error) {
	var latestAnswer models.Answer
	err := c.DB.Where("key = ? AND deleted = ?", key, false).Last(&latestAnswer).Error
	if err != nil {
		return nil, err
	}
	return &latestAnswer, nil
}
