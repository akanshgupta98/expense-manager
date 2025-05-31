package util

import (
	"fmt"
	"net/http"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ReadJSON(c *gin.Context, data any) error {
	err := c.ShouldBindJSON(data)
	if err != nil {
		return err
	}
	logger.Debugf("HTTP payload recieved: %+v", data)
	return nil
}

func WriteJSON(c *gin.Context, data []byte, status ...int) {
	statusCode := http.StatusOK

	if len(status) != 0 {
		statusCode = status[0]
	}
	logger.Debugf("HTTP response sent: %+v with Status: %d", data, statusCode)
	c.Status(statusCode)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Write(data)
}

func ErrorJSON(c *gin.Context, err error, data any, status ...int) {

	statusCode := http.StatusInternalServerError
	var errMsgs []string
	if len(status) != 0 {
		statusCode = status[0]

	}

	errors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, e := range errors {
			errMsg := fmt.Sprintf("Key %s is %s", e.Field(), e.Tag())
			errMsgs = append(errMsgs, errMsg)
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	response := ErrorResponse{
		Err:  errMsgs,
		Data: data,
	}

	logger.Debugf("HTTP Error written: %+v with Status: %d", response, statusCode)

	c.JSON(statusCode, response)

}
