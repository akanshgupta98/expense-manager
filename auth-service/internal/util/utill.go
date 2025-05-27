package util

import (
	"net/http"

	logger "github.com/akanshgupta98/go-logger"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ReadJSON(c *gin.Context, data any) error {

	err := c.ShouldBindJSON(&data)
	if err != nil {
		return err
	}
	logger.Debugf("data in ReadJSON: %+v", data)
	return nil
}

func WriteJSON(c *gin.Context, data any, status ...int) {
	statusCode := http.StatusOK
	if len(status) != 0 {
		statusCode = status[0]
	}
	logger.Debugf("Data written on HTTP: %+v", data)
	c.JSON(statusCode, data)

}

func ErrorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusInternalServerError
	messages := make(map[string]string)
	if len(status) != 0 {
		statusCode = status[0]
	}
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		messages = make(map[string]string)
		for _, e := range errs {
			messages[e.Field()] = e.Tag()
		}
	} else {
		messages["Message"] = err.Error()

	}

	data := ErrorJSONRespone{
		Error: true,
		Msgs:  messages,
	}
	c.JSON(statusCode, data)
}
