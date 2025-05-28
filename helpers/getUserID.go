package helpers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserID mengembalikan user_id dari context dalam bentuk *uint
func GetUserID(c *gin.Context) (*uint, error) {
	userIDAny, exists := c.Get("user_id")
	if !exists {
		return nil, errors.New("user_id not found in context")
	}

	switch v := userIDAny.(type) {
	case uint:
		return &v, nil
	case int:
		val := uint(v)
		return &val, nil
	case int64:
		val := uint(v)
		return &val, nil
	case float64:
		val := uint(v)
		return &val, nil
	case string:
		id64, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, errors.New("failed to parse user_id from string")
		}
		val := uint(id64)
		return &val, nil
	default:
		return nil, errors.New("user_id has unsupported type")
	}
}
