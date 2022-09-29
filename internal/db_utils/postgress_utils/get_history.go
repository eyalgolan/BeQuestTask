package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
	"github.com/pkg/errors"
)

func (c *Client) GetHistory(key string) ([]models.Event, error) {
	var keyHistory []models.Event
	result := c.DB.Where("key = ?", key)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "get history")
	}
	err := result.Find(&keyHistory).Error
	if err != nil {
		return nil, errors.Wrap(err, "get history")
	}
	return keyHistory, nil
}
