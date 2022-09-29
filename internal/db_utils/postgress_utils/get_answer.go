package postgress_utils

import (
	"github.com/eyalgolan/key-value-persistent-store/internal/db_utils/postgress_utils/models"
)

func (c *Client) GetAnswer(key string) (*models.Answer, error) {
	return c.GetLatestAnswer(key)
}
