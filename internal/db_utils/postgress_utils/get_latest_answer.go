package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
)

func (c *Client) GetLatestAnswer(key string) (*models.Answer, error) {
	var latestAnswer models.Answer
	err := c.DB.Where("key = ? AND deleted = ?", key, false).Last(&latestAnswer).Error
	if err != nil {
		return nil, err
	}
	return &latestAnswer, nil
}
