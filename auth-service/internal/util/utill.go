package util

import "github.com/gin-gonic/gin"

func ReadJSON(c *gin.Context, data any) error {

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}
	// logger.Info
	return nil

}
