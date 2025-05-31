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
	logger.Debugf("Data read from HTTP Body is: %+v", data)
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

	var errors []string
	statusCode := http.StatusInternalServerError
	if len(status) != 0 {
		statusCode = status[0]
	}

	errs, ok := err.(validator.ValidationErrors)
	if ok {
		logger.Debugf("Validator errors are: %v", errs)
		for _, e := range errs {
			msg := fmt.Sprintf("%s is %s", e.Field(), errMsg(e.Tag()))
			errors = append(errors, msg)
		}
	} else {
		msg := "unable to perform requested operation."
		errors = append(errors, msg)

	}
	response := Error{
		Errors: errors,
	}

	c.JSON(statusCode, response)
}

func errMsg(tag string) string {
	var msg string

	switch tag {
	case "email":
		msg = "not of type email"

	default:
		msg = tag
	}

	return msg

}
