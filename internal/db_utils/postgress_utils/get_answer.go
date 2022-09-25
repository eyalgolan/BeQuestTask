package postgress_utils

import (
	"KeyValuePermStore/internal/db_utils/postgress_utils/models"
)

func (c *Client) GetAnswer(key string) (*models.Answer, error) {
	return c.getLatestAnswer(key)
}
