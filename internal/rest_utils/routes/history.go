package routes

import (
	"KeyValuePermStore/internal/rest_utils"
	"KeyValuePermStore/internal/rest_utils/gin_context"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
)

func GetHistory(c *gin.Context) {
	db := gin_context.GetDBFromContext(c)
	key := c.Param("key")
	if key == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "GET request must include answer key",
		})
		return
	}
	history, err := db.GetHistory(key)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Not found",
			})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "unable to perform request",
			})
		}
		return
	}

	var keyHistory []rest_utils.Answer

	for _, event := range history {
		keyHistory = append(keyHistory, rest_utils.Answer{
			Event: event.Event,
			Data: rest_utils.AnswerData{
				Key:   event.Key,
				Value: event.Value,
			},
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": keyHistory,
	})
}
