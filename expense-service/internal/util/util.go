package util

import (
	"net/http"

	"github.com/akanshgupta98/go-logger/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ReadJSON(c *gin.Context, response any) error {

	err := c.ShouldBindBodyWithJSON(&response)
	if err != nil {
		return err
	}
	return nil

}

func WriteJSON(c *gin.Context, data any, status ...int) {
	statusCode := http.StatusOK
	if len(status) != 0 {
		statusCode = status[0]
	}
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
	logger.Errorf("Error is: %s", messages)
	c.JSON(statusCode, data)
}
